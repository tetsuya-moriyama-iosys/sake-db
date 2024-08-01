import { createRouter, createWebHistory } from 'vue-router';
import Index from '@/views/IndexView.vue';
import AuthRouter from '@/router/auth';

const subRoutes = [AuthRouter];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    ...subRoutes,
    {
      path: '/',
      name: 'indexView',
      component: Index,
    },
  ],
});

export default router;
