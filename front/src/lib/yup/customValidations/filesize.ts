// ファイルサイズのバリデーション関数
import { mixed } from 'yup';

export const fileSize = (maxSizeInMB: number) => {
  return mixed().test(
    'fileSize',
    `ファイルサイズは${maxSizeInMB}MB以下にしてください`,
    (value: unknown) => {
      if (!value || !(value instanceof File)) {
        return true;
      }

      const maxSizeInBytes = maxSizeInMB * 1024 * 1024;
      return value.size <= maxSizeInBytes;
    },
  );
};
