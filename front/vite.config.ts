import { fileURLToPath, URL } from 'node:url';

import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import { defineConfig } from 'vite';
import graphql from 'vite-plugin-graphql-loader';
import vueDevTools from 'vite-plugin-vue-devtools';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), vueJsx(), vueDevTools(), graphql()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    port: 5173, // フロントエンドのポート
    proxy: {
      '/api': {
        target: 'http://localhost:8080', // バックエンドのポート
        changeOrigin: true, // クロスオリジンヘッダーを正しく設定
        secure: false, // HTTPSでない場合はfalseに
        cookieDomainRewrite: 'localhost', // クッキーのドメインをローカルホストに書き換え
      },
    },
  },
});
