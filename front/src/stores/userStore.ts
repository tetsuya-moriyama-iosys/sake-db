import { defineStore } from 'pinia';
import {
  GET_USER,
  type LoginResponse,
  type LoginResult,
} from '@/graphQL/Auth/auth';
import type { AuthUser } from '@/graphQL/User/user';
import useQuery from '@/funcs/composable/useQuery';

export const USER_STORE = 'user_store';

export const useUserStore = defineStore({
  id: USER_STORE, //ストアの識別子(ロジックでは使わない)
  state: () => ({
    isLogin: false as boolean,
    user: null as AuthUser | null,
  }),
  actions: {
    setUserData(response: LoginResult) {
      console.log('トークンをセット：', response);
      console.log('token：', response.token);
      console.log('user：', response.user);
      localStorage.setItem(import.meta.env.VITE_JWT_TOKEN_NAME, response.token); //ローカルストレージにtokenをセット
      this.isLogin = true; //ログイン状態をtrueにする
      this.user = response.user;
    },
    logout() {
      //ページ遷移はrouterを使って行うため、ストアで実行不可。あくまでも状態のみを変える。
      localStorage.removeItem(import.meta.env.VITE_JWT_TOKEN_NAME);
      this.isLogin = false;
    },
    //画面リロード時などにユーザーデータを取得するために使用
    async restoreUserData() {
      const { fetch } = useQuery<LoginResponse>(GET_USER, {
        isAuth: true,
      });
      const token: string | null = localStorage.getItem(
        import.meta.env.VITE_JWT_TOKEN_NAME,
      );
      // トークンがあり、かつユーザー情報がセットされていない場合は終了
      if (token == null || this.user != null) return;
      try {
        const response: LoginResponse = await fetch(); // APIからユーザー情報を取得
        this.setUserData(response.login); // ユーザー情報をセット
      } catch (error) {
        console.error('ユーザー情報の取得に失敗しました', error);
        this.logout(); // エラー時はログアウト処理
      }
    },
  },
});
