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
      <MenuItem class="block px-4 py-2 text-gray-800 hover:bg-gray-100">
        <router-link :to="{ name: 'Register' }"> 新規登録</router-link>
      </MenuItem>
    </MenuItems>
  </Menu>
</template>

<script setup lang="ts">
import { Menu, MenuButton, MenuItems, MenuItem } from '@headlessui/vue';
</script>

<style scoped></style>
