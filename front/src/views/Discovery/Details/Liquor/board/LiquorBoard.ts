/**
 * カテゴリのフォーム
 */
import * as yup from 'yup';
import { string } from 'yup';

import yupLocaleJP from '@/lib/yup/yupLocaleJa';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  TEXT: 'text', //本文
  RATE: 'rate', //評価
} as const;

export interface FormValues {
  [FormKeys.TEXT]: string;
  [FormKeys.RATE]: number;
}
export const initialValues: FormValues = {
  [FormKeys.TEXT]: '',
  [FormKeys.RATE]: 0,
};

export const validationSchema = {
  [FormKeys.TEXT]: string().min(1).max(500).required(), //とりあえず500文字制限にしておく
};
