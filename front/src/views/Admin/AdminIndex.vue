<template>
  <div>
    <h1>管理者ページ</h1>
    <p>管理者ページです</p>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';

import useQuery from '@/funcs/composable/useQuery';
import { AdminCheck } from '@/graphQL/Admin/admin';
import type { AdminCheckQuery } from '@/graphQL/auto-generated';

const { fetch } = useQuery<AdminCheckQuery>(AdminCheck, {
  isAuth: true,
});
const router = useRouter();

onMounted(async () => {
  await fetch().catch(() => {
    // ログインページへリダイレクト
    router.push({ name: 'Login' });
  });
});
</script>
