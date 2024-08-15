/**
 * 新規作成フォーム
 */
import * as yup from 'yup';
import { string } from 'yup';
import yupLocaleJP from '@/lib/yup/yupLocaleJa';
import { image } from '@/forms/customValidations/image';
import { fileSize } from '@/forms/customValidations/filesize';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  CATEGORY: 'category', //メインカテゴリ
  TITLE: 'title', //名前
  DESCRIPTION: 'description', //説明
  IMAGE: 'image',
} as const;

export interface FormValues {
  [FormKeys.CATEGORY]: number;
  [FormKeys.TITLE]: string;
  [FormKeys.DESCRIPTION]: string;
  [FormKeys.IMAGE]: File | null;
}

export const initialValues = {
  [FormKeys.TITLE]: '',
  [FormKeys.CATEGORY]: null,
  [FormKeys.DESCRIPTION]: '',
  [FormKeys.IMAGE]: null,
};

export const validationSchema = {
  [FormKeys.TITLE]: string().max(100).required(),
  [FormKeys.IMAGE]: image().concat(fileSize(5)).nullable(),
};
