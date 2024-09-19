<template>
  認証済ユーザートップページ <MyPageIndex v-if="user" :user="user" />
</template>

<script setup lang="ts">
import MyPageIndex from '@/components/templates/auth/myPage/MyPageIndex.vue';
import { useUser } from '@/funcs/composable/useUser';

const { user } = useUser();
</script>

<style scoped></style>
