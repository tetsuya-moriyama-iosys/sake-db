import { fileURLToPath, URL } from 'node:url';

import vue from '@vitejs/plugin-vue';
import vueJsx from '@vitejs/plugin-vue-jsx';
import { constants } from 'crypto';
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
    https: {
      key: process.env.VITE_SSL_KEY_PATH,
      cert: process.env.VITE_SSL_CERT_PATH,
      secureOptions: constants.SSL_OP_NO_TLSv1 | constants.SSL_OP_NO_TLSv1_1,
    },
    port: 5173,
  },
});
