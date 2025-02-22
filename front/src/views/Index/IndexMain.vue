<!--メイン画面レイアウト-->
<template>
  <RecommendList v-if="userStore.user" />
  ランダムリスト
  <RandomSelect />
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/userStore/userStore';
import RandomSelect from '@/views/Index/RandomSelect.vue';
import RecommendList from '@/views/Index/RecommendList.vue';

const userStore = useUserStore();
</script>

<style scoped></style>
