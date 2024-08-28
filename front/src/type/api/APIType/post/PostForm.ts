import type { APIType } from '@/type/api/APIType/APIType';
import createURL from '@/funcs/util/core/createURL';

export interface PostRequest {
  category: number;
  title: string;
  description: string;
  image: File | null;
}

export type PostResponse = null;

const PostAPIType: APIType<PostRequest, PostResponse> = {
  method: 'POST',
  url: createURL('post'),
  //ファイル送信はContent-Typeをmultipart/form-dataにする
  headers: {
    'Content-Type': 'multipart/form-data',
  },
};

export default PostAPIType;
