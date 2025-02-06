import { fileURLToPath, URL } from 'node:url';

import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import dotenv from 'dotenv';
import { defineConfig } from 'vite';
import graphql from 'vite-plugin-graphql-loader';
import vueDevTools from 'vite-plugin-vue-devtools';

dotenv.config(); // .env ファイルをロード

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), vueJsx(), vueDevTools(), graphql()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    host: true, // 外部からのアクセスを許可
    port: 5173, // デフォルトの開発サーバーポート
    strictPort: true,
    watch: {
      usePolling: true,
    },
    hmr: {
      host: 'localhost',
      protocol: 'ws', // HTTPS を使用しない場合は `ws`
      port: 5173,
    },
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
