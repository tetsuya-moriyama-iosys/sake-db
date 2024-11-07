/**
 * 新規作成フォーム
 */
import * as yup from 'yup';
import { number, string } from 'yup';

import type { LiquorForEdit } from '@/graphQL/Liquor/liquor';
import { fileSize } from '@/lib/yup/customValidations/filesize';
import { image } from '@/lib/yup/customValidations/image';
import { youtube } from '@/lib/yup/customValidations/youtube';
import yupLocaleJP from '@/lib/yup/yupLocaleJa';
import type { PostRequest } from '@/type/api/APIType/post/PostForm';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  ID: 'id',
  CATEGORY: 'category', //メインカテゴリ
  NAME: 'name', //名前
  DESCRIPTION: 'description', //説明
  YOUTUBE: 'youtube',
  IMAGE: 'image',
  VERSION_NO: 'version_no',
  SELECTED_VERSION_NO: 'selected_version_no',
} as const;

export interface FormValues
  extends Omit<PostRequest, typeof FormKeys.CATEGORY> {
  [FormKeys.CATEGORY]: number | null;
}

const initialValues: FormValues = {
  [FormKeys.ID]: null,
  [FormKeys.NAME]: '',
  [FormKeys.CATEGORY]: null,
  [FormKeys.DESCRIPTION]: '',
  [FormKeys.YOUTUBE]: '',
  [FormKeys.IMAGE]: null,
  [FormKeys.SELECTED_VERSION_NO]: null,
  [FormKeys.VERSION_NO]: null,
};

export function generateInitialValues(
  liquor: LiquorForEdit | null,
): FormValues {
  if (liquor === null) return initialValues;
  return {
    ...initialValues, //imageなどは編集時も空なので初期値を設定
    [FormKeys.ID]: liquor.id,
    [FormKeys.NAME]: liquor.name,
    [FormKeys.CATEGORY]: liquor.categoryId,
    [FormKeys.DESCRIPTION]: liquor.description,
    [FormKeys.VERSION_NO]: liquor.versionNo,
  };
}

export const validationSchema = {
  [FormKeys.CATEGORY]: number().required(),
  [FormKeys.NAME]: string().max(100).required(),
  [FormKeys.DESCRIPTION]: string().max(5000),
  [FormKeys.IMAGE]: image().concat(fileSize(2)).nullable(),
  [FormKeys.YOUTUBE]: youtube().nullable(),
};
