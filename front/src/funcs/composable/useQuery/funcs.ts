import type { QueryOptions } from '@apollo/client/core/watchQueryOptions';
import type { DocumentNode } from 'graphql';
import type { FieldNode, OperationDefinitionNode } from 'graphql/index';
import { type Ref, ref, watch } from 'vue';
import {
  type ActiveLoader,
  type PluginApi,
  useLoading,
} from 'vue-loading-overlay';
import { type Router } from 'vue-router';

import type { QueryOption } from '@/funcs/composable/useQuery/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import { errorDebug } from '@/funcs/util/core/console';
import { refreshToken, useUserStore } from '@/stores/userStore/userStore';

const spinner: PluginApi = useLoading();

/**
 * @package
 */
export function getCommon<T>(operation: DocumentNode, option?: QueryOption) {
  const toast = useToast();
  const loading = ref<boolean>(false);
  const data = ref<T>();
  let activeSpinner: ActiveLoader | null = null;

  const operationName: string = (
    (operation.definitions[0] as OperationDefinitionNode).selectionSet
      .selections[0] as FieldNode
  ).name.value;

  const showSpinner = (): ActiveLoader => {
    return spinner.show();
  };

  const hideSpinner = () => {
    if (activeSpinner) {
      activeSpinner.hide();
      activeSpinner = null;
    }
  };

  // watchでloadingの状態を監視し、スピナーの表示・非表示を切り替える
  watch(loading, (newVal) => {
    if (newVal && option?.isUseSpinner === true) {
      // loadingがtrueになったらスピナーを表示
      activeSpinner = showSpinner();
    } else if (!newVal && activeSpinner) {
      // loadingがfalseになったらスピナーを非表示
      hideSpinner();
    }
  });

  const handleError = (err: Error) => {
    errorDebug('handleError: ', err);
    errorDebug('エラー：', err.message);
    toast.errorToast((err.message as string) || '不明なエラー');
  };

  return { operationName, loading, data, handleError };
}

/**
 * @package
 */
export function generate<V = unknown>(
  request: V,
  options?: Omit<QueryOptions, 'query' | 'variables'>,
  customOptions?: QueryOption,
) {
  //渡されたオプションからヘッダーを生成
  const headers = options?.context?.headers || {};

  // isAuthフラグがtrueの場合、JWTトークンを追加
  if (customOptions?.isAuth) {
    // ストアにアクセスしてアクセストークンを取得
    const { accessToken } = useUserStore();
    const token = accessToken.get();
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }
  }

  //リクエストをvariablesでラップ
  const variables = request ? { variables: request } : {};

  return {
    variables,
    headers,
  };
}

// クエリ実行関数(アクセストークンが切れたら再チャレンジする)
interface ChallengeOption<T> {
  run: () => Promise<T>;
  operationName: string;
  data: Ref<T | undefined, T | undefined>;
  loading: Ref<boolean>;
  router: Router; //injectがsetup内でしか呼び出せないので、引数で渡す(グローバルオブジェクトを直接インポートしても可能らしいが、やめておく)
}
export async function challenge<T = unknown>({
  run,
  operationName,
  data,
  loading,
  router,
}: ChallengeOption<T>): Promise<T> {
  loading.value = true;

  try {
    const response = await run();
    data.value = response;
    return response;
  } catch (err: unknown) {
    errorDebug('エラー返却：', err);
    if ((err as Error).message == 'token expired') {
      await refreshToken(); //アクセストークン期限切れの場合、リフレッシュトークンを再取得

      //再度リクエスト
      const response = await run();
      data.value = response;
      return response;
    } else if ((err as Error).message == 'unauthorized') {
      //認証エラーの場合はログインページにリダイレクト
      void router.push({ name: 'Login' });
    } else {
      errorDebug('error at ', operationName);
    }
    throw err; // 親にエラーを投げる
  } finally {
    loading.value = false;
  }
}
