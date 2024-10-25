import 'vue-loading-overlay/dist/css/index.css';

import { VueQueryPlugin } from '@tanstack/vue-query';
import { createPinia } from 'pinia';
import { type App } from 'vue';
import { LoadingPlugin } from 'vue-loading-overlay';

import ToastPlugin from '@/plugins/toast';

import router from './router';

export function registerPlugins(app: App) {
  /**ライブラリ*/
  // Pinia
  app.use(createPinia());

  // TanStack Query
  app.use(VueQueryPlugin);

  // Vue Router
  app.use(router);

  //スピナー
  app.use(LoadingPlugin);

  /**カスタムプラグイン*/
  // トーストプラグイン
  app.use(ToastPlugin);
}
