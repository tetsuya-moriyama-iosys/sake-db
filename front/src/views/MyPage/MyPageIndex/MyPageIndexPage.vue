<template>
  認証済ユーザートップページ <MyPageIndex v-if="user" :user="user" />
</template>

<script setup lang="ts">
import { useUserDetail } from '@/views/MyPage/funcs';
import MyPageIndex from '@/views/MyPage/MyPageIndex/MyPageIndex.vue';

const { user } = useUserDetail();
</script>

<style scoped></style>
