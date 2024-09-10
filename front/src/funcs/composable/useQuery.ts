/**
 * GraphQL用クライアント
 */
import { ref, watch } from 'vue';
import {
  type DocumentNode,
  type FieldNode,
  type OperationDefinitionNode,
} from 'graphql';
import client from '@/apolloClient';
import type {
  MutationOptions,
  QueryOptions,
} from '@apollo/client/core/watchQueryOptions';
import type { ApolloQueryResult } from '@apollo/client/core/types';
import { useToast } from '@/funcs/composable/useToast';
import type { FetchResult } from '@apollo/client';
import {
  type ActiveLoader,
  type PluginApi,
  useLoading,
} from 'vue-loading-overlay';

const spinner: PluginApi = useLoading();

export interface QueryOption {
  isUseSpinner?: boolean; //読み込み時のスピナーを表示する
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

  const handleError = (err: unknown) => {
    error.value = err;
    console.error('エラー：', err);
    toast.errorToast((err as string) || '不明なエラー');
  };

  return { loading, error, data, handleError };
}

export function useQuery<T = unknown>(
  query: DocumentNode,
  option?: QueryOption,
) {
  const { loading, error, data, handleError } = useCommon<T>(option);

  const fetch = async (options?: Omit<QueryOptions, 'query'>): Promise<T> => {
    loading.value = true;
    error.value = null;
    try {
      const result: ApolloQueryResult<T> = await client.query<T>({
        ...options,
        query,
      });
      //型推論がうまくいってないのでキャスト
      console.log(
        (
          (query.definitions[0] as OperationDefinitionNode).selectionSet
            .selections[0] as FieldNode
        ).name.value,
        'レスポンス：',
        result.data,
      );
      data.value = result.data;
    } catch (err) {
      handleError(err);
      throw err;
    } finally {
      loading.value = false;
    }

    return data.value;
  };

  return { fetch, loading, error, data };
}

export function useMutation<T = unknown>(
  mutation: DocumentNode,
  option?: QueryOption,
) {
  const { loading, error, data, handleError } = useCommon<T>(option);

  const execute = async (
    options?: Omit<MutationOptions<T>, 'mutation'>,
  ): Promise<T> => {
    loading.value = true;
    error.value = null;
    try {
      const result: FetchResult<T> = await client.mutate<T>({
        ...options,
        mutation,
      });
      console.log(
        (
          (mutation.definitions[0] as OperationDefinitionNode).selectionSet
            .selections[0] as FieldNode
        ).name.value,
        'レスポンス：',
        result.data,
      );
      data.value = result.data as T; //ライブラリのジェネリクスがundefinedも含んでいるのでキャスト。返す場合は大元のジェネリクスの方で指定する。
    } catch (err) {
      handleError(err);
      throw err;
    } finally {
      loading.value = false;
    }

    return data.value;
  };

  return { execute, loading, error, data };
}

export default useQuery;
