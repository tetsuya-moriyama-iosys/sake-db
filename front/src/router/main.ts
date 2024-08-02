/**
 * メイン画面(データベース検索部分)のルーター
 */
import Layout from '@/views/Discovery/Layouts/BaseLayout.vue';
import IndexPage from '@/views/Discovery/IndexPage.vue';

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
  ],
};

export default MainRouter;
