<template>
  <header class="sticky">
    <main class="flex">
      <p class="block">酒データベース(仮称)</p>
      <router-link to="/post">投稿する</router-link>
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
