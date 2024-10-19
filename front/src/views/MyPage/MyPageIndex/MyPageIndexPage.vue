<template>
  認証済ユーザートップページ <MyPageIndex v-if="user" :user="user" />
</template>

<script setup lang="ts">
import { useUser } from '@/funcs/composable/useUser';
import MyPageIndex from '@/views/MyPage/MyPageIndex/MyPageIndex.vue';

const { user } = useUser();
</script>

<style scoped></style>
