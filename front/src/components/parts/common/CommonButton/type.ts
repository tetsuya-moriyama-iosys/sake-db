import { ColorType } from '@/type/common/ColorType';

export interface ButtonProps {
  color?: ColorType;
  size?: 'small' | 'large';
  class?: string;
  isDisabled?: boolean;
}
