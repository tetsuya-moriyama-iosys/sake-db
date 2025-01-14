<template>
  ユーザー情報編集
  <UserEditForm v-if="user" :user="user" />
</template>

<script setup lang="ts">
import { useUserDetail } from '@/views/MyPage/funcs';
import UserEditForm from '@/views/MyPage/UserEdit/form/UserEditForm.vue';

const { user } = useUserDetail();
</script>

<style scoped></style>
