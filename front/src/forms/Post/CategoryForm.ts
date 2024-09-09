/**
 * カテゴリのフォーム
 */
import * as yup from 'yup';
import { number, string } from 'yup';
import yupLocaleJP from '@/lib/yup/yupLocaleJa';
import { image } from '@/forms/customValidations/image';
import { fileSize } from '@/forms/customValidations/filesize';
import type { CategoryRequest } from '@/type/api/APIType/post/CategoryForm';
import type { Category } from '@/graphQL/Category/categories';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  ID: 'id',
  PARENT: 'parent', //メインカテゴリ
  NAME: 'name', //名前
  DESCRIPTION: 'description', //説明
  IMAGE: 'image',
  SELECTED_VERSION_NO: 'selected_version_no',
  VERSION_NO: 'version_no',
} as const;

export interface FormValues
  extends Omit<CategoryRequest, typeof FormKeys.PARENT> {
  [FormKeys.PARENT]: number | null;
}

export const initialValues: FormValues = {
  [FormKeys.ID]: null,
  [FormKeys.NAME]: '',
  [FormKeys.PARENT]: null,
  [FormKeys.DESCRIPTION]: '',
  [FormKeys.IMAGE]: null,
  [FormKeys.SELECTED_VERSION_NO]: null,
  [FormKeys.VERSION_NO]: null,
};

export function generateInitialValues(category: Category | null): FormValues {
  if (category === null) {
    return initialValues;
  }

  const value = {
    ...initialValues, //imageなどは編集時も空なので初期値を設定
    [FormKeys.ID]: category.id,
    [FormKeys.NAME]: category.name,
    [FormKeys.PARENT]: category.parent,
    [FormKeys.DESCRIPTION]: category.description,
    [FormKeys.SELECTED_VERSION_NO]: category.versionNo,
  };

  console.log('generateInitialValues:', value);
  return value;
}

export const validationSchema = {
  [FormKeys.PARENT]: number().required(),
  [FormKeys.NAME]: string().max(100).required(),
  [FormKeys.IMAGE]: image().concat(fileSize(2)).nullable(),
};
