/**
 * GraphQL用クライアント
 */
import { ref } from 'vue';
import { type DocumentNode } from 'graphql';
import client from '@/apolloClient';

export function useQuery<T = unknown>(query: DocumentNode) {
  const loading = ref<boolean>(false);
  const error = ref<unknown>(null);
  const data = ref<T>();

  const fetch = async (): Promise<T> => {
    loading.value = true;
    error.value = null;
    try {
      const result = await client.query({
        query,
      });
      data.value = result.data;
    } catch (err) {
      error.value = err;
      throw err;
    } finally {
      loading.value = false;
    }

    return data.value as T;
  };

  return { fetch, loading, error, data };
}

export default useQuery;
