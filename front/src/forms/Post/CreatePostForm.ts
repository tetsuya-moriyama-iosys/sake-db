/**
 * 新規作成フォーム
 */
import * as yup from 'yup';
import { string } from 'yup';
import yupLocaleJP from '@/lib/yup/yupLocaleJa';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  CATEGORY: 'category', //メインカテゴリ
  TITLE: 'title', //名前
  DESCRIPTION: 'description', //説明
} as const;

export const initialValues = {
  [FormKeys.TITLE]: '初期タイトル',
  [FormKeys.CATEGORY]: 11,
};

export const validationSchema = {
  [FormKeys.TITLE]: string().max(5).required(),
};
