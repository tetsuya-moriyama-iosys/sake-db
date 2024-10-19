<template>
  ユーザー情報編集
  <UserEditForm v-if="user" :user="user" />
</template>

<script setup lang="ts">
import { useUser } from '@/funcs/composable/useUser';
import UserEditForm from '@/views/MyPage/UserEdit/form/UserEditForm.vue';

const { user } = useUser();
</script>

<style scoped></style>
