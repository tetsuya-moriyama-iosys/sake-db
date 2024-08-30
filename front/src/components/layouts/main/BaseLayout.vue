<template>
  <div>
    <HeaderArea />
    <main class="flex">
      <SideBar />
      <div class="flex-1">
        <router-view></router-view>
      </div>
    </main>
    <FooterArea />
  </div>
</template>
<script setup lang="ts">
import HeaderArea from '@/components/layouts/main/HeaderArea.vue';
import FooterArea from '@/components/layouts/main/FooterArea.vue';
import SideBar from '@/components/layouts/main/SideBar.vue';
</script>

<style scoped>
div {
  height: calc(100vh - 100px);
  width: 100%;

  main {
    height: calc(100% - 100px); /* ヘッダーとフッターの高さを差し引いた高さ */
  }
}
</style>
