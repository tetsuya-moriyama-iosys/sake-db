import { ColorType } from '@/type/common/ColorType';

export interface TagProps {
  text: string;
  isClose?: boolean;
  color?: ColorType;
  size?: 'small' | 'large';
}
