/**
 * ユーザー情報関連のルーター
 */
import UserIndex from '@/views/User/UserIndex.vue';

export const MyPageRouter = {
  path: 'my-page',
  meta: { requiresAuth: true },
  children: [
    {
      path: 'index',
      name: 'UserIndex',
      component: UserIndex,
    },
  ],
};
