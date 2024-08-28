import { inject } from 'vue';
import { TOAST_INJECT_KEY, type ToastCommand } from '@/plugins/toast';

export const useToast = (): ToastCommand => {
  const toast = inject<ToastCommand>(TOAST_INJECT_KEY);
  if (!toast) {
    throw new Error('Toast plugin is not properly injected');
  }
  return toast;
};
