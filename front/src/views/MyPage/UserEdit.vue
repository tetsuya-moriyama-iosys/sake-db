<template>
  ユーザー情報編集
  <UserEditForm v-if="user" :user="user" />
</template>
<script setup lang="ts">
import { useUserFullData } from '@/funcs/composable/useUser';
import UserEditForm from '@/components/templates/auth/myPage/UserEditForm.vue';

const { user } = useUserFullData();
</script>

<style scoped></style>
