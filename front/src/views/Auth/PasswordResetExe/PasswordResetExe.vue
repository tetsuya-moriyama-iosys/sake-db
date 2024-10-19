<!--パスワードリセットの実行-->
<template>
  パスワードリセット実行
  <PasswordResetExeForm :token="token" />
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';

import PasswordResetExeForm from '@/views/Auth/PasswordResetExe/form/PasswordResetExeForm.vue';

const route = useRoute(); // 現在のルートを取得
const token: string = route.params.token as string;
</script>

<style scoped></style>
