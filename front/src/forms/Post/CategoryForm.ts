/**
 * カテゴリのフォーム
 */
import * as yup from 'yup';
import { number, string } from 'yup';
import yupLocaleJP from '@/lib/yup/yupLocaleJa';
import { image } from '@/forms/customValidations/image';
import { fileSize } from '@/forms/customValidations/filesize';
import type { CategoryRequest } from '@/type/api/APIType/post/CategoryForm';
import type { CategoryForEdit } from '@/graphQL/Liquor/categories';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  ID: 'id',
  PARENT: 'parent', //メインカテゴリ
  NAME: 'name', //名前
  DESCRIPTION: 'description', //説明
  IMAGE: 'image',
  VERSION_NO: 'version_no',
} as const;

export interface FormValues
  extends Omit<CategoryRequest, typeof FormKeys.PARENT> {
  [FormKeys.PARENT]: number | null;
}

const initialValues: FormValues = {
  [FormKeys.ID]: null,
  [FormKeys.NAME]: '',
  [FormKeys.PARENT]: null,
  [FormKeys.DESCRIPTION]: '',
  [FormKeys.IMAGE]: null,
  [FormKeys.VERSION_NO]: null,
};

export function generateInitialValues(
  category: CategoryForEdit | null,
): FormValues {
  if (category === null) return initialValues;
  return {
    ...initialValues, //imageなどは編集時も空なので初期値を設定
    [FormKeys.ID]: category.id,
    [FormKeys.NAME]: category.name,
    [FormKeys.PARENT]: category.parent,
    [FormKeys.DESCRIPTION]: category.description,
    [FormKeys.VERSION_NO]: category.versionNo,
  };
}

export const validationSchema = {
  [FormKeys.PARENT]: number().required(),
  [FormKeys.NAME]: string().max(100).required(),
  [FormKeys.IMAGE]: image().concat(fileSize(2)).nullable(),
};
