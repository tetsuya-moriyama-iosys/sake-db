/**
 * 登録フォーム
 */
import * as yup from 'yup';
import { string } from 'yup';
import yupLocaleJP from '@/lib/yup/yupLocaleJa';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  NAME: 'name',
  MAIL: 'email',
  PASSWORD: 'password',
} as const;

export interface FormValues {
  [FormKeys.NAME]: string;
  [FormKeys.MAIL]: string;
  [FormKeys.PASSWORD]: string;
}

export const initialValues: FormValues = {
  [FormKeys.NAME]: '',
  [FormKeys.MAIL]: '',
  [FormKeys.PASSWORD]: '',
};

export const validationSchema = {
  [FormKeys.NAME]: string().max(50).required(),
  [FormKeys.MAIL]: string().email().required(),
  [FormKeys.PASSWORD]: string().min(7).required(),
};
