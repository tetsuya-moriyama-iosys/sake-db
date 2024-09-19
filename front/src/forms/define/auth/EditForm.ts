/**
 * 登録情報編集フォーム
 */
import * as yup from 'yup';
import yupLocaleJP from '@/lib/yup/yupLocaleJa';
import {
  FormKeys as RegisterFormKeys,
  type FormValues as RegisterFormValues,
  validationSchema as RegisterValidationSchema,
} from './RegisterForm';
import { string } from 'yup';
import type { UserFullData } from '@/graphQL/User/user';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  ...RegisterFormKeys,
  IMAGE: 'image',
  PROFILE: 'profile',
} as const;

export interface FormValues extends RegisterFormValues {
  [FormKeys.IMAGE]: File | null;
  [FormKeys.PROFILE]: string;
}

export function generateInitialValues(user: UserFullData): FormValues {
  return {
    [FormKeys.NAME]: user.name,
    [FormKeys.MAIL]: user.email,
    [FormKeys.PASSWORD]: '', //パスワードは入力が必須ではない
    [FormKeys.IMAGE]: null,
    [FormKeys.PROFILE]: user.profile,
  };
}
export const validationSchema = {
  ...RegisterValidationSchema,
  [FormKeys.PASSWORD]: string()
    .transform((value) => (value === '' ? null : value))
    .nullable()
    .min(7),
};
