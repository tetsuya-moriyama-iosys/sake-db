<template>
  <header class="sticky">
    <main class="flex">
      <router-link to="/"
        ><p class="block">酒データベース(仮称)</p></router-link
      >
      <router-link :to="{ name: 'LiquorEdit' }">投稿する</router-link>
      <router-link :to="{ name: 'CategoryEdit' }">新規カテゴリ追加</router-link>
    </main>
  </header>
</template>

<script setup lang="ts"></script>

<style scoped>
header {
  width: 100%;
  height: 100px;
  background-color: black;
  color: white;
}
</style>
