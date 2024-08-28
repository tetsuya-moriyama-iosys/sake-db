/**
 * トーストのプラグイン
 */
/* eslint-disable no-unused-vars */
import { type App, createApp, h, ref } from 'vue';
import CommonToast from '@/components/parts/common/CommonToast.vue';

export const TOAST_INJECT_KEY = Symbol('toast');

export enum ToastType {
  Success = 'success',
  Error = 'error',
  Info = 'info',
}

export interface ToastOptions {
  message: string;
  type?: ToastType;
  duration?: number;
}

export interface ToastCommand {
  showToast: (options: ToastOptions) => void;
  errorToast: (errorMsg: string) => void;
}

const DEFAULT_TOAST_STATUS = {
  message: '',
  type: ToastType.Success,
  duration: 2000,
};

export const toastState = ref<ToastOptions | null>(null);

const showToast = (options: ToastOptions) => {
  toastState.value = {
    message: options.message,
    type: options.type ?? DEFAULT_TOAST_STATUS.type,
    duration: options.duration ?? DEFAULT_TOAST_STATUS.duration,
  };
};

const createToast = (): ToastCommand => {
  const app = createApp({
    render: () => h(CommonToast, toastState.value),
  });
  const mountNode = document.createElement('div');
  document.body.appendChild(mountNode);
  app.mount(mountNode);

  return {
    showToast,
    errorToast,
  };
};

export const errorToast = (errorMsg: string) => {
  showToast({ message: errorMsg, type: ToastType.Error });
};

export default {
  install(app: App) {
    const toast = createToast();
    app.provide<ToastCommand>(TOAST_INJECT_KEY, toast);
    app.config.globalProperties.$toast = toast;
  },
};
