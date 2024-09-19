<template>ユーザーデータ 名前：{{ props.user.name }}</template>
<script setup lang="ts">
import type { User } from '@/graphQL/User/user';

interface Props {
  user: User;
}

const props = defineProps<Props>();
</script>
<style scoped></style>
