import { defineStore } from 'pinia';
import { nextTick, ref } from 'vue';

import useQuery from '@/funcs/composable/useQuery';
import {
  type AuthUser,
  GET_USER,
  type GetUserResponse,
  type LoginResult,
  type Role,
} from '@/graphQL/Auth/auth';

export const USER_STORE = 'user_store';

export const useUserStore = defineStore(USER_STORE, () => {
  const { fetch } = useQuery<GetUserResponse>(GET_USER, {
    isAuth: true,
  });

  const isLogin = ref<boolean>(false);
  const user = ref<AuthUser | null>(null);

  function setUserData(response: LoginResult) {
    localStorage.setItem(import.meta.env.VITE_JWT_TOKEN_NAME, response.token); //ローカルストレージにtokenをセット
    isLogin.value = true; //ログイン状態をtrueにする

    user.value = {
      id: response.user.id,
      name: response.user.name,
      imageBase64: response.user.imageBase64,
      roles: response.user.roles,
    };
  }

  //画面リロード時・情報アップデート時などにユーザーデータを取得するために使用(情報を変えない限りキャッシュを使った方がいい)
  async function restoreUserData(option?: { isReFetch?: boolean }) {
    const token: string | null = localStorage.getItem(
      import.meta.env.VITE_JWT_TOKEN_NAME,
    );

    // トークンがない場合は終了
    if (token == null) {
      return;
    }
    try {
      // APIからユーザー情報を取得(ユーザー情報)
      const response: GetUserResponse = await fetch(
        undefined,
        option?.isReFetch === true
          ? {
              fetchPolicy: 'no-cache',
            }
          : {},
      );
      setUserData({
        token: token, //ログイン時とインターフェースを合わせるために追加。//TODO: リフレッシュトークンの実装
        user: response.getUser, // ユーザー情報をセット
      });

      // nextTickでUIの更新を保証
      await nextTick();
    } catch (error) {
      console.error('ユーザー情報の取得に失敗しました', error);
      logout(); // エラー時はログアウト処理
    }
  }

  function getRoles(): Role[] {
    return user.value?.roles ?? [];
  }

  function logout() {
    //ページ遷移はrouterを使って行うため、ストアで実行不可。あくまでも状態のみを変える。
    localStorage.removeItem(import.meta.env.VITE_JWT_TOKEN_NAME);
    //ストア情報のクリア
    isLogin.value = false;
    user.value = null;
  }

  return { isLogin, user, setUserData, logout, getRoles, restoreUserData };
});
