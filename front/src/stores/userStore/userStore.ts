import { defineStore } from 'pinia';
import { nextTick, ref } from 'vue';

import useQuery from '@/funcs/composable/useQuery/useQuery';
import {
  type AuthUser,
  GET_USER,
  type GetUserResponse,
  type Role,
} from '@/graphQL/Auth/auth';
import { type AuthPayloadForUI } from '@/stores/userStore/type';

export const USER_STORE = 'user_store';

export const useUserStore = defineStore(USER_STORE, () => {
  const { fetch } = useQuery<GetUserResponse>(GET_USER, {
    isAuth: true,
  });

  const isLogin = ref<boolean>(false);
  const user = ref<AuthUser | null>(null);
  const accessToken = ref<string | null>(null);

  //ログイン情報をストアにセットする
  function setUserData(data: AuthPayloadForUI) {
    accessToken.value = data.accessToken;
    isLogin.value = true; //ログイン状態をtrueにする
    user.value = data.user; //ユーザー情報をセット
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
        accessToken: token, //ログイン時とインターフェースを合わせるために追加。
        user: response.getUser, // ユーザー情報をセット
      });

      await nextTick(); // nextTickでUIの更新を保証
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
