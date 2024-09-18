/**
 * ユーザー情報関連のルーター
 */
import UserPage from '@/views/User/UserPage.vue';

export const UserRouter = {
  path: 'user/:id',
  children: [
    {
      path: 'index',
      name: 'UserPage',
      component: UserPage,
    },
  ],
};
