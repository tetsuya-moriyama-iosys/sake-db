<template>
  <RouterView />
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router';

import { useUserStore } from '@/stores/userStore/userStore';
const userStore = useUserStore();

// アプリケーションがロードされた時に、ストアのユーザー情報を復元する
userStore.restoreUserData();
</script>

<style scoped></style>
