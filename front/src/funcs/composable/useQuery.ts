/**
 * GraphQL用クライアント
 */
import { ref } from 'vue';
import {
  type DocumentNode,
  type FieldNode,
  type OperationDefinitionNode,
} from 'graphql';
import client from '@/apolloClient';
import type { QueryOptions } from '@apollo/client/core/watchQueryOptions';
import type { ApolloQueryResult } from '@apollo/client/core/types';
import { useToast } from '@/funcs/composable/useToast';

export function useQuery<T = unknown>(query: DocumentNode) {
  const toast = useToast();
  const loading = ref<boolean>(false);
  const error = ref<unknown>(null);
  const data = ref<T>();

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
      error.value = err;
      //共通のエラートースト処理
      console.error('エラー：', err);
      toast.errorToast((err as string) || '不明なエラー');
      throw err;
    } finally {
      loading.value = false;
    }

    return data.value;
  };

  return { fetch, loading, error, data };
}

export default useQuery;
