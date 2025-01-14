import { createRouter, createWebHistory } from 'vue-router';

import MainRouter from '@/router/main';
import { authenticate } from '@/router/middleware/authenticate';
import MetaComponent from '@/views/MetaInfo.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'MetaView',
      component: MetaComponent, // 親ルート用の空コンポーネント
      children: [MainRouter],
    },
  ],
});

//認証ミドルウェア
router.beforeEach((to, from, next) => {
  authenticate(to, next);
  next();
});

export default router;
