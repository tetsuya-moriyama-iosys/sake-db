import { defineStore } from 'pinia';

export const USER_STORE = 'user_store';

// Base64デコード用の関数
function decodeBase64Url(base64: string): string {
  const base64String = base64.replace(/-/g, '+').replace(/_/g, '/');
  const decodedString = atob(base64String);

  // Convert decoded string into a Uint8Array and then use TextDecoder
  const uint8Array = new Uint8Array(
    [...decodedString].map((char) => char.charCodeAt(0)),
  );
  const textDecoder = new TextDecoder('utf-8');
  return textDecoder.decode(uint8Array);
}

function generateUser(): User {
  const token: string = localStorage.getItem(
    import.meta.env.VITE_JWT_TOKEN_NAME,
  ) as string; //ここにアクセスできる時点で存在することは確定している

  // JWTトークンからペイロードを取り出し、ユーザーIDを取得
  const payloadBase64 = token.split('.')[1]; // トークンのペイロード部分
  const payloadJson = decodeBase64Url(payloadBase64);
  const payload = JSON.parse(payloadJson);

  console.log('トークン内情報', payload);

  return { id: payload.id };
}

export type User = {
  id: string;
};

export const useUserStore = defineStore({
  id: USER_STORE,
  state: () => ({
    isLogin: false as boolean,
    user: null as User | null,
  }),
  actions: {
    setToken(token: string) {
      localStorage.setItem(import.meta.env.VITE_JWT_TOKEN_NAME, token); //ローカルストレージにtokenをセット
      this.isLogin = true; //ログイン状態をtrueにする
      this.user = generateUser();
    },
    logout() {
      //ページ遷移はrouterを使って行うため、ストアで実行不可。あくまでも状態のみを変える。
      localStorage.removeItem(import.meta.env.VITE_JWT_TOKEN_NAME);
      this.isLogin = false;
    },
  },
});
