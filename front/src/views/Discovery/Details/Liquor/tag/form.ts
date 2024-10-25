import { string } from 'yup';

import { type PostTag, PostTagKeys } from '@/graphQL/Liquor/tags';

export type FormValues = PostTag;

export function defaultValues(liquorId: string): FormValues {
  return {
    [PostTagKeys.LiquorId]: liquorId,
    [PostTagKeys.Tag]: '',
  };
}

export const validationSchema = {
  [PostTagKeys.Tag]: string().min(1).max(20).required(),
};
