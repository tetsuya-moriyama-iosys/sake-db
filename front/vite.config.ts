import { fileURLToPath, URL } from 'node:url';

import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import dotenv from 'dotenv';
import { defineConfig } from 'vite';
import vueDevTools from 'vite-plugin-vue-devtools';

dotenv.config(); // .env ファイルをロード

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), vueJsx(), vueDevTools()],
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
    // proxy: {
    //   '/api': {
    //     target: 'https://localhost',
    //     changeOrigin: true,
    //     secure: true,
    //   },
    // },
  },
  // server: {
  //   https: {
  //     key: process.env.VITE_SSL_KEY_PATH,
  //     cert: process.env.VITE_SSL_CERT_PATH,
  //     secureOptions: constants.SSL_OP_NO_TLSv1 | constants.SSL_OP_NO_TLSv1_1,
  //   },
  //   port: 5173,
  // },
});
