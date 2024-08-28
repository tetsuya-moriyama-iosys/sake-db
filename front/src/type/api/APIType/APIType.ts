import type { MaybeRefDeep } from '@tanstack/vue-query/build/legacy/types';
import type { MutationObserver } from 'vue-query';
import type { AxiosResponse } from 'axios';

export interface APIType<
  Request extends object | null = object,
  Response extends object | null = object,
> {
  url: string;
  method: 'POST' | 'GET';
  headers?: {
    'Content-Type': ContentType;
    [key: string]: string; // その他のヘッダーも含められるようにする
  };
  options?: MaybeRefDeep<
    MutationObserver<AxiosResponse<Response>, unknown, Request>
  >;
}

export type ContentType =
  | 'application/json'
  | 'multipart/form-data'
  | 'application/x-www-form-urlencoded'
  | 'text/plain';
