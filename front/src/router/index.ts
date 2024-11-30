import { createRouter, createWebHistory } from 'vue-router';

import MainRouter from '@/router/main';
import MetaComponent from '@/views/MetaInfo.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'MetaView',
      component: MetaComponent, // 親ルート用の空コンポーネント
      children: [MainRouter],
    },
  ],
});

//認証ミドルウェア
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth === undefined) {
    //認証がそもそも不要なページの場合はそのまま遷移
    next();
    return;
  }
  const token: string | null = localStorage.getItem(
    import.meta.env.VITE_JWT_TOKEN_NAME,
  );
  if (token != null) {
    // トークンをデコードしてペイロードを取得
    const tokenPayload = JSON.parse(atob(token.split('.')[1]));

    // トークンの有効期限を確認
    const currentTime: number = Math.floor(Date.now() / 1000);
    if (tokenPayload.exp && tokenPayload.exp > currentTime) {
      // ここで有効なトークンとして扱い、APIリクエストなどを行う
      next();
      return;
    } else {
      // トークンが無効・有効期限切れの場合はトークンを削除し、再ログインを促す
      localStorage.removeItem(import.meta.env.VITE_JWT_TOKEN_NAME);
    }
  }
  // 認証が必要なページにアクセスしようとしたが、トークンがない場合はログインページにリダイレクト
  next({ name: 'Login' });
});

export default router;
