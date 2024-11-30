/**
 * メイン画面(データベース検索部分)のルーター
 */
import Layout from '@/components/layouts/main/BaseLayout.vue';
import AdminRouter from '@/router/admin';
import AuthRouter from '@/router/auth';
import { MyPageRouter } from '@/router/mypage';
import { UserRouter } from '@/router/user';
import CategoryPage from '@/views/Discovery/Details/Category/CategoryPage.vue';
import DetailPage from '@/views/Discovery/Details/Liquor/LiquorPage.vue';
import CategoryNarrowDownPage from '@/views/Discovery/NarrowDowns/CategoryNarrowDownPage.vue';
import CategoryEditPage from '@/views/Edit/CategoryEdit/CategoryEditPage.vue';
import CreatePost from '@/views/Edit/LiquorEdit/LiquorEditPage.vue';
import IndexPage from '@/views/Index/IndexPage.vue';

const MainRouter = {
  path: '/',
  name: 'Layout',
  component: Layout,
  children: [
    //Layoutを共用するため、この位置に移動
    AuthRouter,
    MyPageRouter,
    UserRouter,
    AdminRouter,
    {
      path: '/',
      name: 'Index',
      component: IndexPage,
    },

    {
      path: '/discovery',
      children: [
        {
          path: 'category/:id?',
          name: 'CategoryNarrowDown',
          component: CategoryNarrowDownPage,
        },
      ],
    },
    {
      path: '/liquor',
      children: [
        { path: ':id?', name: 'LiquorDetail', component: DetailPage },
        {
          path: 'edit/:id?',
          name: 'LiquorEdit',
          component: CreatePost,
        },
      ],
    },
    {
      path: '/category',
      children: [
        { path: ':id', name: 'CategoryDetail', component: CategoryPage },
        {
          path: 'edit/:id?',
          name: 'CategoryEdit',
          component: CategoryEditPage,
        },
      ],
    },
  ],
};

export default MainRouter;
