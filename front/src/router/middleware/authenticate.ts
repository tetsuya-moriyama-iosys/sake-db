//認証ミドルウェア
import type {
  NavigationGuardNext,
  RouteLocationNormalizedGeneric,
} from 'vue-router';

import { useUserStore } from '@/stores/userStore/userStore';

export function authenticate(
  to: RouteLocationNormalizedGeneric,
  next: NavigationGuardNext,
) {
  if (to.meta.requiresAuth === undefined) {
    //認証がそもそも不要なページの場合はそのまま遷移
    next(); // NOTE: nextを呼び忘れるとミドルウェアの実行が止まるので注意
    return;
  }

  const userStore = useUserStore();
  console.log('ユーザーストア：', userStore);
  const token = userStore.accessToken.get();

  if (token == null) {
    console.error('認証失敗！');
    // 認証が必要なページにアクセスしようとしたが、トークンがない場合はログインページにリダイレクト
    next({ name: 'Login' });
    return;
  }
  console.log('認証通過：');
  //NOTE: 実際にトークンを使用する際にリフレッシュトークンを使って再発行する処理を挟むので、ミドルウェアで破棄する必要はない。あくまでもトークンの有無のみ確認すればOKになった。

  // トークンをデコードしてペイロードを取得
  //const tokenPayload = JSON.parse(atob(token.split('.')[1]));

  // トークンの有効期限を確認
  // const currentTime: number = Math.floor(Date.now() / 1000);
  // if (tokenPayload.exp && tokenPayload.exp > currentTime) {
  //   // ここで有効なトークンとして扱い、APIリクエストなどを行う
  //   next();
  //   return;
  // } else {
  //   // トークンが無効・有効期限切れの場合はトークンを削除し、再ログインを促す
  //   localStorage.removeItem(import.meta.env.VITE_JWT_TOKEN_NAME);
  // }
  next();
}
