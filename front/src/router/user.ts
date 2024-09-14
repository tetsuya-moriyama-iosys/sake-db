/**
 * ユーザー情報関連のルーター
 */
import UserIndex from '@/views/User/UserIndex.vue';

export const UserRouter = {
  path: 'user/:id',
  meta: { requiresAuth: true },
  children: [
    {
      path: 'index',
      name: 'UserIndex',
      component: UserIndex,
    },
  ],
};
