import { type App } from 'vue';
import { createPinia } from 'pinia';
import ToastPlugin from '@/plugins/toast';
import { VueQueryPlugin } from '@tanstack/vue-query';
import router from './router';

export function registerPlugins(app: App) {
  /**ライブラリ*/
  // Pinia
  app.use(createPinia());

  // TanStack Query
  app.use(VueQueryPlugin);

  // Vue Router
  app.use(router);

  /**カスタムプラグイン*/
  // トーストプラグイン
  app.use(ToastPlugin);
}
