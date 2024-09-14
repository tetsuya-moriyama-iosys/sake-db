<template>認証済ユーザートップページ{{ user }}</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useUserStore } from '@/stores/userStore';
import useQuery from '@/funcs/composable/useQuery';
import { GET_USERDATA, type GetUserdataResponse } from '@/graphQL/User/user';
import type { User } from '@/graphQL/Auth/auth';

const userStore = useUserStore();

const { fetch } = useQuery<GetUserdataResponse>(GET_USERDATA, {
  isAuth: true,
});

const userData = ref<User | null>(null);

// Base64デコード用の関数
function decodeBase64Url(base64: string) {
  const base64String = base64.replace(/-/g, '+').replace(/_/g, '/');
  const decodedString = atob(base64String);

  // Convert decoded string into a Uint8Array and then use TextDecoder
  const uint8Array = new Uint8Array(
    [...decodedString].map((char) => char.charCodeAt(0)),
  );
  const textDecoder = new TextDecoder('utf-8');
  return textDecoder.decode(uint8Array);
}

const user = computed(() => {
  const token: string = localStorage.getItem(
    import.meta.env.VITE_JWT_TOKEN_NAME,
  ) as string; //ここにアクセスできる時点で存在することは確定している

  // JWTトークンからペイロードを取り出し、ユーザーIDを取得
  const payloadBase64 = token.split('.')[1]; // トークンのペイロード部分
  const payloadJson = decodeBase64Url(payloadBase64);
  const payload = JSON.parse(payloadJson);

  console.log('トークン内情報', payload);

  return payload;
});

onMounted(async () => {
  userStore.checkAuthentication(); //同タブのstorage変更はキャッチできないので、手動で呼び出す
  const response: GetUserdataResponse = await fetch();
  userData.value = response.getUser;
});
</script>

<style scoped></style>
