/**
 * GraphQL用クライアント
 */
import { ref } from 'vue';
import { type DocumentNode } from 'graphql';
import client from '@/apolloClient';
import type { QueryOptions } from '@apollo/client/core/watchQueryOptions';
import type { ApolloQueryResult } from '@apollo/client/core/types';

export function useQuery<T = unknown>(query: DocumentNode) {
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
      data.value = result.data;
    } catch (err) {
      error.value = err;
      throw err;
    } finally {
      loading.value = false;
    }

    return data.value;
  };

  return { fetch, loading, error, data };
}

export default useQuery;
