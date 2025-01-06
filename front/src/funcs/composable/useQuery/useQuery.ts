/**
 * GraphQL用クライアント
 */
import type { FetchResult } from '@apollo/client';
import type { ApolloQueryResult } from '@apollo/client/core/types';
import type {
  MutationOptions,
  QueryOptions,
} from '@apollo/client/core/watchQueryOptions';
import {
  type DocumentNode,
  type FieldNode,
  type OperationDefinitionNode,
} from 'graphql';
import { ref, watch } from 'vue';
import {
  type ActiveLoader,
  type PluginApi,
  useLoading,
} from 'vue-loading-overlay';
import { type Router, useRouter } from 'vue-router';

import client from '@/apolloClient';
import { useToast } from '@/funcs/composable/useToast';

const spinner: PluginApi = useLoading();

export interface QueryOption {
  isUseSpinner?: boolean; //読み込み時のスピナーを表示する
  isAuth?: boolean; // JWT認証が必要かどうか
}

function useCommon<T>(option?: QueryOption) {
  const toast = useToast();
  const loading = ref<boolean>(false);
  const error = ref<unknown>(null);
  const data = ref<T>();
  let activeSpinner: ActiveLoader | null = null;

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
    error.value = err.message;
    console.error('エラー：', error.value);
    toast.errorToast((error.value as string) || '不明なエラー');
  };

  return { loading, error, data, handleError };
}

function generate<V = unknown>(
  request: V,
  options?: Omit<QueryOptions, 'query' | 'variables'>,
  customOptions?: QueryOption,
) {
  //渡されたオプションからヘッダーを生成
  const headers = options?.context?.headers || {};

  // isAuthフラグがtrueの場合、JWTトークンを追加
  if (customOptions?.isAuth) {
    const token = localStorage.getItem(import.meta.env.VITE_JWT_TOKEN_NAME);
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

export function useQuery<T = unknown, V = unknown>(
  query: DocumentNode,
  option?: QueryOption,
) {
  const { loading, error, data, handleError } = useCommon<T>(option);
  const router: Router = useRouter();
  const queryName: string = (
    (query.definitions[0] as OperationDefinitionNode).selectionSet
      .selections[0] as FieldNode
  ).name.value;

  const fetch = async (
    request?: V,
    options?: Omit<QueryOptions, 'query' | 'variables'>,
  ): Promise<T> => {
    loading.value = true;
    error.value = null;

    try {
      const { variables, headers } = generate(request, options, option);

      console.log(queryName, '送信データ：', variables);
      //TODO: なんか急に型エラーが出始めたので調査の必要がありそう
      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      //@ts-expect-error
      const result: ApolloQueryResult<T> = await client.query<T>({
        ...variables,
        ...options,
        query,
        context: {
          headers, // ヘッダーをセット
        },
      });
      //型推論がうまくいってないのでキャスト
      console.log(queryName, 'レスポンス：', result.data);
      data.value = result.data;
    } catch (err: unknown) {
      console.error('error at ', queryName);
      if ((err as Error).message == 'unauthorized') {
        //認証エラーの場合はログインページにリダイレクト
        void router.push({ name: 'Login' });
      } else {
        handleError(err as Error);
      }
      throw err;
    } finally {
      loading.value = false;
    }

    return data.value;
  };

  return { fetch, loading, error, data };
}

export function useMutation<T = unknown, V = unknown>(
  mutation: DocumentNode,
  option?: QueryOption,
) {
  const { loading, error, data, handleError } = useCommon<T>(option);
  const mutationName: string = (
    (mutation.definitions[0] as OperationDefinitionNode).selectionSet
      .selections[0] as FieldNode
  ).name.value;

  const execute = async (
    request: V,
    options?: Omit<MutationOptions<T>, 'mutation'>,
  ): Promise<T> => {
    loading.value = true;
    error.value = null;
    try {
      const { variables, headers } = generate(request, options, option);

      console.log(mutationName, '送信データ：', variables);
      //TODO: なんか急に型エラーが出始めたので調査の必要がありそう
      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      //@ts-expect-error
      const result: FetchResult<T> = await client.mutate<T>({
        ...variables,
        ...options,
        mutation,
        context: {
          headers, // ヘッダーをセット
        },
      });
      console.log(mutationName, 'レスポンス：', result.data);
      data.value = result.data as T; //ライブラリのジェネリクスがundefinedも含んでいるのでキャスト。返す場合は大元のジェネリクスの方で指定する。
    } catch (err) {
      console.error('error at ', mutationName);
      handleError(err as Error);
      throw err;
    } finally {
      loading.value = false;
    }

    return data.value;
  };

  return { execute, loading, error, data };
}

export default useQuery;
