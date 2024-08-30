/**
 * メイン画面(データベース検索部分)のルーター
 */
import Layout from '@/components/layouts/main/BaseLayout.vue';
import IndexPage from '@/views/Discovery/IndexPage.vue';
import CreatePost from '@/views/Post/CreatePost.vue';
import DetailPage from '@/views/Discovery/Liquor/DetailPage.vue';

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
      path: '/post',
      name: 'Post',
      component: CreatePost,
    },
    {
      path: '/liquor/:id',
      name: 'LiquorDetail',
      component: DetailPage,
    },
  ],
};

export default MainRouter;
