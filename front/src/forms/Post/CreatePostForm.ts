/**
 * 新規作成フォーム
 */
import * as yup from 'yup';
import { number, string } from 'yup';
import yupLocaleJP from '@/lib/yup/yupLocaleJa';
import { image } from '@/forms/customValidations/image';
import { fileSize } from '@/forms/customValidations/filesize';
import type { PostRequest } from '@/type/api/APIType/post/PostForm';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  CATEGORY: 'category', //メインカテゴリ
  TITLE: 'title', //名前
  DESCRIPTION: 'description', //説明
  IMAGE: 'image',
} as const;

export type FormValues = PostRequest;

export const initialValues = {
  [FormKeys.TITLE]: '',
  [FormKeys.CATEGORY]: null,
  [FormKeys.DESCRIPTION]: '',
  [FormKeys.IMAGE]: null,
};

export const validationSchema = {
  [FormKeys.CATEGORY]: number().required(),
  [FormKeys.TITLE]: string().max(100).required(),
  [FormKeys.IMAGE]: image().concat(fileSize(5)).nullable(),
};
