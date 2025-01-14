<template>
  ログイン
  <LoginForm />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';

import { useUserStore } from '@/stores/userStore/userStore';
import LoginForm from '@/views/Auth/Login/form/LoginForm.vue';

const { isLogin } = useUserStore();
const router = useRouter();

//すでにログイン済の場合、ユーザーページへリダイレクト
onMounted(() => {
  if (isLogin) {
    router.push({ name: 'UserIndex' });
  }
});
</script>

<style scoped></style>
