<template>
  ログイン
  <LoginForm />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';

import LoginForm from '@/components/templates/auth/login/LoginForm.vue';

// const { isLogin } = useLogin();
// const router = useRouter();

//すでにログイン済の場合、ユーザーページへリダイレクト
onMounted(() => {
  // TODO:IDがないらしいので、取得or保存する処理が必要
  // if (isLogin.value) {
  //   router.push({ name: 'UserIndex' });
  // }
});
</script>

<style scoped></style>
