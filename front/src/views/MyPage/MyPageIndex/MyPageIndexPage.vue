<template>
  認証済ユーザートップページ <MyPageIndex v-if="user" :user="user" />
</template>

<script setup lang="ts">
import { getUserDetail } from '@/views/MyPage/funcs';
import MyPageIndex from '@/views/MyPage/MyPageIndex/MyPageIndex.vue';

const { user } = getUserDetail();
</script>

<style scoped></style>
