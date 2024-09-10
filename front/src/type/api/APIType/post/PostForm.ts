import type { APIType } from '@/type/api/APIType/APIType';
import createURL from '@/funcs/util/core/createURL';

export interface PostRequest {
  id: string | null;
  category: number;
  name: string;
  description: string;
  image: File | null;
  selected_version_no: number | null;
  version_no: number | null;
}

export type PostResponse = {
  id: string;
};

const PostAPIType: APIType<PostRequest, PostResponse> = {
  method: 'POST',
  url: createURL('post'),
  //ファイル送信はContent-Typeをmultipart/form-dataにする
  headers: {
    'Content-Type': 'multipart/form-data',
  },
};

export default PostAPIType;
