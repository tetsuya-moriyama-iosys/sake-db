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
        <router-link :to="{ name: 'Register' }"
          ><div @click="close()">新規登録</div></router-link
        >
      </MenuItem>
      <MenuItem
        v-if="!userStore.isLogin"
        v-slot="{ close }"
        class="block px-4 py-2 text-gray-800 hover:bg-gray-100"
      >
        <router-link :to="{ name: 'Login' }"
          ><div @click="close()">ログイン</div></router-link
        >
      </MenuItem>
      <div v-else>
        <MenuItem
          class="block px-4 py-2 text-gray-800 hover:bg-gray-100"
          v-slot="{ close }"
        >
          <router-link :to="{ name: 'MyPageIndex' }"
            ><div @click="close()">マイページ</div></router-link
          >
        </MenuItem>
        <MenuItem
          class="block px-4 py-2 text-gray-800 hover:bg-gray-100"
          v-slot="{ close }"
        >
          <div @click="handleLogout(close)">ログアウト</div>
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
