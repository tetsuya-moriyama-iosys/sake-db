import { type App } from 'vue';
import { createPinia } from 'pinia';
import ToastPlugin from '@/plugins/toast';
import { VueQueryPlugin } from '@tanstack/vue-query';
import router from './router';
import { LoadingPlugin } from 'vue-loading-overlay';
import 'vue-loading-overlay/dist/css/index.css';

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
