<template>
  <div>
    <HeaderArea />
    <main class="flex">
      <section class="side-bar fixed">
        <SideBar />
      </section>
      <section class="main-container overflow-y-scroll">
        <router-view></router-view>
      </section>
    </main>
    <FooterArea />
  </div>
</template>
<script setup lang="ts">
import HeaderArea from '@/components/layouts/main/header/HeaderArea.vue';
import FooterArea from '@/components/layouts/main/FooterArea.vue';
import SideBar from '@/components/layouts/main/sideBar/SideBar.vue';
</script>

<style scoped>
div {
  height: calc(100vh - 150px);
  width: 100%;

  main {
    section {
      height: calc(
        100vh - 150px
      ); /* ヘッダーとフッターの高さを差し引いた高さ */
      &.main-container {
        margin-left: 180px;
        width: calc(100% - 180px);
      }
    }
  }
}
</style>
