<!--メイン画面レイアウト-->
<template>
  <RecommendList v-if="userStore.user" />
  ランダムリスト
  <RandomSelect />
</template>

<script setup lang="ts">
import RandomSelect from '@/components/templates/index/RandomSelect.vue';
import { useUserStore } from '@/stores/userStore';
import RecommendList from '@/components/templates/index/RecommendList.vue';

const userStore = useUserStore();
</script>

<style scoped></style>
