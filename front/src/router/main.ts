/**
 * メイン画面(データベース検索部分)のルーター
 */
import Layout from '@/components/layouts/main/BaseLayout.vue';
import IndexPage from '@/views/Discovery/IndexPage.vue';
import CreatePost from '@/views/Post/LiquorEditPage.vue';
import DetailPage from '@/views/Discovery/Liquor/LiquorPage.vue';
import CategoryPage from '@/views/Discovery/Category/CategoryPage.vue';
import CategoryEditPage from '@/views/Post/CategoryEditPage.vue';

const MainRouter = {
  path: '/',
  name: 'Layout',
  component: Layout,
  children: [
    {
      path: '/',
      name: 'Index',
      component: IndexPage,
    },
    {
      path: '/post/:id?',
      name: 'LiquorEdit',
      component: CreatePost,
    },
    {
      path: '/liquor',
      children: [{ path: ':id?', name: 'LiquorDetail', component: DetailPage }],
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
