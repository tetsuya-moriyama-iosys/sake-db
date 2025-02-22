<template>
  <header class="sticky">
    <main class="flex">
      <section class="flex flex-1">
        <router-link to="/"
          ><p class="block">酒データベース(仮称)</p></router-link
        >
        <router-link :to="{ name: 'LiquorEdit' }">投稿する</router-link>
      </section>
      <section v-if="getRoles().includes(Roles.Admin)">
        <router-link :to="{ name: 'Admin' }">管理者ページ</router-link>
      </section>
      <section>
        <AccountInfo />
      </section>
      <section>
        <MainMenu />
      </section>
    </main>
  </header>
</template>

<script setup lang="ts">
import AccountInfo from '@/components/layouts/main/header/AccountInfo.vue';
import MainMenu from '@/components/layouts/main/header/menu/MainMenu.vue';
import { Roles } from '@/graphQL/Auth/types';
import { useUserStore } from '@/stores/userStore/userStore';

const { getRoles } = useUserStore();
</script>

<style scoped>
header {
  width: 100%;
  height: 100px;
  background-color: black;
  color: white;
}
</style>
