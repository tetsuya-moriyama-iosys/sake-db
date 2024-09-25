<template>
  名前：{{ props.user.name }}
  <router-link :to="{ name: 'MyPageEdit' }">ユーザー情報編集</router-link>
  <div>
    <BookmarkList />
  </div>
</template>

<script setup lang="ts">
import type { AuthUser } from '@/graphQL/Auth/auth';
import BookmarkList from '@/components/templates/auth/myPage/BookmarkList.vue';

interface Props {
  user: AuthUser;
}

const props = defineProps<Props>();
</script>

<style scoped></style>
