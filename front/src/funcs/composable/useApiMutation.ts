import { useMutation } from '@tanstack/vue-query';
import axios, { type AxiosResponse } from 'axios';
import { unref } from 'vue';

import { useToast } from '@/funcs/composable/useToast';
import { debug } from '@/funcs/util/core/console';
import type { APIType } from '@/type/api/APIType/APIType';

export function useApiMutation<
  Request extends object | null = object,
  Response extends object | null = object,
>(
  apiType: APIType<Request, Response>,
  options: { isAuth: boolean } = { isAuth: false },
) {
  const toast = useToast();
  const mutationFn = async (
    data: Request,
  ): Promise<AxiosResponse<Response>> => {
    debug('リクエストdata:', data);
    const header = apiType.headers ?? { 'Content-Type': 'application/json' };

    async function run(): Promise<AxiosResponse<Response>> {
      return axios({
        url: apiType.url,
        method: apiType.method,
        headers: options.isAuth ? { ...header } : header,
        data,
      }).then((response) => {
        debug('リクエスト成功 - レスポンス:', response.data);
        return response;
      });
    }

    return run().catch((error) => {
      //共通のエラートースト処理
      toast.errorToast(error?.response?.data?.error || '不明なエラー');
      throw error;
    });
  };

  return useMutation<AxiosResponse<Response>, unknown, Request>({
    mutationFn,
    ...unref(apiType.options), // optionsがMaybeRefDeepである可能性があるため、unrefで取り出す
  });
}
