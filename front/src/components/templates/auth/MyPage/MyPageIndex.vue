<template>
  名前：{{ props.user.name }}
  <router-link :to="{ name: 'MyPageEdit' }">ユーザー情報編集</router-link>
</template>

<script setup lang="ts">
import type { AuthUser } from '@/graphQL/User/user';

interface Props {
  user: AuthUser;
}

const props = defineProps<Props>();
</script>

<style scoped></style>
