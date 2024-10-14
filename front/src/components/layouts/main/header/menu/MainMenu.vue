<template>
  <Menu>
    <MenuButton>
      <svg
        class="w-6 h-6"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 6h16M4 12h16M4 18h16"
        ></path>
      </svg>
    </MenuButton>
    <MenuItems
      class="absolute right-0 mt-2 py-2 w-48 bg-white rounded-md shadow-xl"
    >
      <MenuItem link-to="Register">新規登録</MenuItem>
      <MenuItem v-if="!userStore.isLogin" link-to="Login">ログイン</MenuItem>
      <div v-else>
        <MenuItem link-to="MyPageIndex">マイページ</MenuItem>
        <MenuItem link-to="Index" @click="logout">ログアウト</MenuItem>
      </div>
    </MenuItems>
  </Menu>
</template>

<script setup lang="ts">
import { Menu, MenuButton, MenuItems } from '@headlessui/vue';

import { useToast } from '@/funcs/composable/useToast';
import { useUserStore } from '@/stores/userStore';

import MenuItem from './MenuItem.vue';

const userStore = useUserStore();
const toast = useToast();

function logout() {
  userStore.logout();
  toast.showToast({
    message: 'ログアウトしました',
  });
}
</script>

<style scoped></style>
