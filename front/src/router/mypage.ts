/**
 * ユーザー情報関連のルーター
 */
import UserIndex from '@/views/MyPage/MyPageIndex.vue';
import UserEdit from '@/views/MyPage/UserEdit.vue';

export const MyPageRouter = {
  path: 'my-page',
  meta: { requiresAuth: true },
  children: [
    {
      path: 'index',
      name: 'MyPageIndex',
      component: UserIndex,
    },
    {
      path: 'edit',
      name: 'MyPageEdit',
      component: UserEdit,
    },
  ],
};