/**
 * GraphQL用クライアント
 * NOTE: 返しているrefオブジェクト、便利かなと思って入れたけど使わなさそうなので消してもいいかも知れない
 *       ちなみにmutationとqueryで引数が微妙に異なっており、強引な共通化をしようとすると型管理がややこしくなったのでやめた
 */
import type { FetchResult } from '@apollo/client';
import type { ApolloQueryResult } from '@apollo/client/core/types';
import type {
  MutationOptions,
  QueryOptions,
} from '@apollo/client/core/watchQueryOptions';
import { type DocumentNode } from 'graphql';
import { useRouter } from 'vue-router';

import client from '@/apolloClient';
import {
  challenge,
  generate,
  getCommon,
} from '@/funcs/composable/useQuery/funcs';
import { debug } from '@/funcs/util/core/console';

export interface QueryOption {
  isUseSpinner?: boolean; //読み込み時のスピナーを表示する
  isAuth?: boolean; // JWT認証が必要かどうか
}

export function useQuery<T = unknown, V = unknown>(
  query: DocumentNode,
  option?: QueryOption,
) {
  const router = useRouter();
  const { operationName, loading, data, handleError } = getCommon<T>(
    query,
    option,
  );

  async function run(
    request?: V,
    options?: Omit<QueryOptions, 'query' | 'variables'>,
  ): Promise<T> {
    const { variables, headers } = generate(request, options, option);

    debug(operationName, '送信データ：', variables);
    //ここで例外が投げられるとfetchでキャッチする
    const result = (await client.query<T>({
      ...variables,
      ...options,
      query,
      context: {
        headers, // ヘッダーをセット
      },
    })) as ApolloQueryResult<T>; // MaybeMaskedを剥がす
    debug(operationName, 'レスポンス：', result.data);

    data.value = result.data;
    return result.data;
  }

  async function fetch(
    request?: V,
    options?: Omit<QueryOptions, 'query' | 'variables'>,
  ): Promise<T> {
    try {
      return challenge<T>({
        run: () => run(request, options),
        operationName: operationName,
        data,
        loading,
        router,
      });
    } catch (e) {
      handleError(e as Error);
      throw e;
    }
  }

  return { fetch, loading, data };
}

export function useMutation<Response = unknown, V = unknown>(
  mutation: DocumentNode,
  option?: QueryOption,
) {
  const router = useRouter();
  const { operationName, loading, data, handleError } = getCommon<Response>(
    mutation,
    option,
  );

  async function run(
    request: V,
    options?: Omit<MutationOptions<Response>, 'mutation'>,
  ): Promise<Response> {
    const { variables, headers } = generate(request, options, option);

    debug(operationName, '送信データ：', variables);
    const result = (await client.mutate<Response>({
      ...variables,
      ...options,
      mutation,
      context: {
        headers, // ヘッダーをセット
      },
    })) as FetchResult<Response>; // MaybeMaskedを剥がす
    debug(operationName, 'レスポンス：', result.data);

    return result.data as Response; //ライブラリのジェネリクスがundefinedも含んでいるのでキャスト。返す場合は大元のジェネリクスの方で指定する。
  }

  async function execute(
    request: V,
    options?: Omit<MutationOptions<Response>, 'mutation'>,
  ): Promise<Response> {
    try {
      return challenge<Response>({
        run: () => run(request, options),
        operationName: operationName,
        data,
        loading,
        router,
      });
    } catch (e) {
      handleError(e as Error);
      throw e;
    }
  }

  return { execute, loading, data };
}

export default useQuery;
