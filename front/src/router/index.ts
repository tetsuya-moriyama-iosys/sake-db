import { createRouter, createWebHistory } from 'vue-router';
import MetaComponent from '@/views/MetaInfo.vue';
import AuthRouter from '@/router/auth';
import MainRouter from '@/router/main';

const routes = [MainRouter, AuthRouter];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'MetaView',
      component: MetaComponent, // 親ルート用の空コンポーネント
      children: [...routes],
    },
  ],
});

export default router;
