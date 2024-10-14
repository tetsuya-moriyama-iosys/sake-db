/**
 * 登録フォーム
 */
import * as yup from 'yup';
import { string } from 'yup';

import yupLocaleJP from '@/lib/yup/yupLocaleJa';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  MAIL: 'email',
  PASSWORD: 'password',
} as const;

export interface FormValues {
  [FormKeys.MAIL]: string;
  [FormKeys.PASSWORD]: string;
}

export const initialValues: FormValues = {
  [FormKeys.MAIL]: '',
  [FormKeys.PASSWORD]: '',
};

export const validationSchema = {
  [FormKeys.MAIL]: string().email().required(),
  [FormKeys.PASSWORD]: string(),
};
