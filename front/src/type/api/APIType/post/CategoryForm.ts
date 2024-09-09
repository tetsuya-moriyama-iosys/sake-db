import type { APIType } from '@/type/api/APIType/APIType';
import createURL from '@/funcs/util/core/createURL';

export interface CategoryRequest {
  id: number | null;
  parent: number;
  name: string;
  description: string;
  image: File | null;
  selected_version_no: number | null;
  version_no: number | null;
}

export type CategoryResponse = {
  id: number;
};

const CategoryPostAPIType: APIType<CategoryRequest, CategoryResponse> = {
  method: 'POST',
  url: createURL('category/post'),
  //ファイル送信はContent-Typeをmultipart/form-dataにする
  headers: {
    'Content-Type': 'multipart/form-data',
  },
};

export default CategoryPostAPIType;
