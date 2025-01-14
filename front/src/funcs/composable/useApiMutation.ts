import { useMutation } from '@tanstack/vue-query';
import axios, { type AxiosResponse } from 'axios';
import { unref } from 'vue';

import { useToast } from '@/funcs/composable/useToast';
import { debug } from '@/funcs/util/core/console';
import type { APIType } from '@/type/api/APIType/APIType';

export function useApiMutation<
  Request extends object | null = object,
  Response extends object | null = object,
>(apiType: APIType<Request, Response>) {
  const toast = useToast();
  const mutationFn = async (
    data: Request,
  ): Promise<AxiosResponse<Response>> => {
    debug('リクエストdata:', data);
    return axios({
      url: apiType.url,
      method: apiType.method,
      headers: apiType.headers ?? { 'Content-Type': 'application/json' },
      data,
    })
      .then((response) => {
        debug('リクエスト成功 - レスポンス:', response.data);
        return response;
      })
      .catch((error) => {
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
