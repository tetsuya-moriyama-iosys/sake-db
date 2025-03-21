import { defineStore } from 'pinia';
import { nextTick, ref } from 'vue';

import client from '@/apolloClient';
import { useMutation } from '@/funcs/composable/useQuery/useQuery';
import { errorDebug } from '@/funcs/util/core/console';
import {
  LOGIN_WITH_REFRESH_TOKEN,
  LOGOUT,
  REFRESH_TOKEN,
} from '@/graphQL/Auth/auth';
import { type AuthUser, type Role } from '@/graphQL/Auth/types';
import type {
  LoginWithRefreshTokenMutation,
  LogoutMutation,
} from '@/graphQL/auto-generated';
import {
  type AuthPayloadForUI,
  getAuthPayloadForUI,
} from '@/stores/userStore/type';

export const USER_STORE = 'user_store';

export const useUserStore = defineStore(USER_STORE, () => {
  const { execute: refreshTokenExecute } =
    useMutation<LoginWithRefreshTokenMutation>(LOGIN_WITH_REFRESH_TOKEN, {
      isAuth: true,
    });
  const { execute: logoutExecute } = useMutation<LogoutMutation>(LOGOUT, {
    isAuth: true,
  });

  const isLogin = ref<boolean>(false);
  const user = ref<AuthUser | null>(null);
  const accessTokenRef = ref<string | null>(null);

  //ログイン情報をストアにセットする
  function setUserData(data: AuthPayloadForUI) {
    accessToken.reset(data.accessToken);
    isLogin.value = true; //ログイン状態をtrueにする
    user.value = data.user; //ユーザー情報をセット
  }

  const accessToken = {
    get: (): string | null => {
      return accessTokenRef.value;
    },
    reset: (newToken: string) => {
      accessTokenRef.value = newToken;
    },
  } as const;

  //画面リロード時・情報アップデート時などにユーザーデータを取得するために使用(情報を変えない限りキャッシュを使った方がいい)
  async function restoreUserData(option?: { isReFetch?: boolean }) {
    console.log('restore userStore');
    try {
      // APIからユーザー情報を取得(ユーザー情報)
      const { loginWithRefreshToken: payload } = await refreshTokenExecute(
        undefined,
        option?.isReFetch === true
          ? {
              fetchPolicy: 'no-cache',
            }
          : {},
      );
      // ユーザーデータをストアに格納する
      setUserData(getAuthPayloadForUI(payload));

      await nextTick(); // nextTickでUIの更新を保証
    } catch (error) {
      errorDebug('restoreUserData失敗', error); // トーストは必要ない
    }
  }

  function getRoles(): Role[] {
    return user.value?.roles ?? [];
  }

  async function logout() {
    //ページ遷移はrouterを使って行うため、ストアで実行不可。あくまでも状態のみを変える。
    await logoutExecute(undefined);
    //ストア情報のクリア
    isLogin.value = false;
    user.value = null;
  }

  return {
    isLogin,
    user,
    setUserData,
    logout,
    getRoles,
    restoreUserData,
    accessToken,
  };
});

export async function refreshToken() {
  const {
    accessToken: { reset },
  } = useUserStore();
  const result = await client.mutate({
    mutation: REFRESH_TOKEN,
    context: {
      credentials: 'include', // クッキーを送信する
    },
  });
  // リフレッシュトークンから取得したアクセストークンをセット
  reset(result.data.refreshToken.accessToken);
}
