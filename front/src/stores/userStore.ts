import { defineStore } from 'pinia';
import { isEmpty } from '@/funcs/util/isEmpty';

export const USER_STORE = 'user_store';
export const useUserStore = defineStore({
  id: USER_STORE,
  state: () => ({
    isLogin: false as boolean,
  }),
  actions: {
    checkAuthentication() {
      const token: string | null = localStorage.getItem(
        import.meta.env.VITE_JWT_TOKEN_NAME,
      );
      this.isLogin = !isEmpty(token); // トークンがあればtrue、なければfalse
    },
    logout() {
      //ページ遷移はrouterを使って行うため、ストアで実行不可。あくまでも状態のみを変える。
      localStorage.removeItem(import.meta.env.VITE_JWT_TOKEN_NAME);
      this.isLogin = false;
    },
  },
});
