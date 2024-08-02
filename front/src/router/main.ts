/**
 * メイン画面(データベース検索部分)のルーター
 */
import Layout from '@/views/Main/Layouts/BaseLayout.vue';
import IndexPage from '@/views/Main/IndexPage.vue';

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
