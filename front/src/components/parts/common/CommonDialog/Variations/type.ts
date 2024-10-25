import type { DialogProps } from '@/components/parts/common/CommonDialog/type';

export type YesNoDialogProps = Omit<DialogProps, 'isUnUseDefaultButtons'> & {
  yes?: string;
  no?: string;
  onYes: () => void;
  onNo?: () => void;
};
