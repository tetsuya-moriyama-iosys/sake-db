/**
 * メイン画面(データベース検索部分)のルーター
 */
import Layout from '@/components/layouts/main/BaseLayout.vue';
import IndexPage from '@/views/Discovery/IndexPage.vue';
import CreatePost from '@/views/Edit/LiquorEditPage.vue';
import DetailPage from '@/views/Discovery/Details/LiquorPage.vue';
import CategoryPage from '@/views/Discovery/Details/CategoryPage.vue';
import CategoryEditPage from '@/views/Edit/CategoryEditPage.vue';
import CategoryNarrowDownPage from '@/views/Discovery/NarrowDowns/CategoryNarrowDownPage.vue';
import AuthRouter from '@/router/auth';
import { UserRouter } from '@/router/user';

const MainRouter = {
  path: '/',
  name: 'Layout',
  component: Layout,
  children: [
    //Layoutを共用するため、この位置に移動
    AuthRouter,
    UserRouter,
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
