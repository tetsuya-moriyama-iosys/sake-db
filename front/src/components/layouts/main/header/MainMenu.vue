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
      <MenuItem
        v-slot="{ close }"
        class="block px-4 py-2 text-gray-800 hover:bg-gray-100"
      >
        <span>
          <router-link :to="{ name: 'Register' }" @click="close()"
            >新規登録</router-link
          ></span
        >
      </MenuItem>
      <MenuItem
        v-if="!userStore.isLogin"
        v-slot="{ close }"
        class="block px-4 py-2 text-gray-800 hover:bg-gray-100"
      >
        <span>
          <router-link :to="{ name: 'Login' }" @click="close"
            >ログイン</router-link
          >
        </span>
      </MenuItem>
      <div v-else>
        <MenuItem
          class="block px-4 py-2 text-gray-800 hover:bg-gray-100"
          v-slot="{ close }"
        >
          <span>
            <router-link :to="{ name: 'MyPageIndex' }" @click="close"
              >マイページ</router-link
            >
          </span>
        </MenuItem>
        <MenuItem
          class="block px-4 py-2 text-gray-800 hover:bg-gray-100"
          v-slot="{ close }"
        >
          <span @click="handleLogout(close)">ログアウト</span>
        </MenuItem>
      </div>
    </MenuItems>
  </Menu>
</template>

<script setup lang="ts">
import { Menu, MenuButton, MenuItems, MenuItem } from '@headlessui/vue';
import { useUserStore } from '@/stores/userStore';
import { useRouter } from 'vue-router';

const userStore = useUserStore();
const router = useRouter();

const handleLogout = (close: () => void) => {
  userStore.logout(); //ログアウト処理を実行する
  close(); //メニューを閉じる
  //routerはコンポーネントからしか呼び出せない
  router.push({ name: 'Index' });
};
</script>

<style scoped></style>
