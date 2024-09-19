<template>
  ユーザー情報編集
  <UserEditForm v-if="user" :user="user" />
</template>
<script setup lang="ts">
import UserEditForm from '@/components/templates/auth/myPage/UserEditForm.vue';
import { useUser } from '@/funcs/composable/useUser';

const { user } = useUser();
</script>

<style scoped></style>
