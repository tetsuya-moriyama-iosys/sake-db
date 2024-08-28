import { mixed, MixedSchema } from 'yup';

export const image = (): MixedSchema => {
  return mixed().test(
    'fileType',
    '画像形式のみアップロード可能です',
    (value: unknown) => {
      if (!value || !(value instanceof File)) {
        return true;
      }

      const validTypes = ['image/jpeg', 'image/png', 'image/gif'];
      return validTypes.includes(value.type);
    },
  );
};
