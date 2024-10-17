<template>
  名前：{{ props.user.name }}
  <router-link :to="{ name: 'MyPageEdit' }">ユーザー情報編集</router-link>
  <div>
    <BookmarkList />
  </div>
</template>

<script setup lang="ts">
import type { AuthUserFull } from '@/graphQL/Auth/auth';
import BookmarkList from '@/views/MyPage/MyPageIndex/BookmarkList.vue';

interface Props {
  user: AuthUserFull;
}

const props = defineProps<Props>();
</script>

<style scoped></style>
