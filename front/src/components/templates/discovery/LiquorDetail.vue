<!--酒詳細ページ-->

<template>
  <p class="title">{{ liquor.name }}</p>
  <img v-if="liquor.imageUrl" :src="liquor.imageUrl" alt="画像" />
  <router-link :to="{ name: 'LiquorEdit', params: { id: liquor.id } }">
    <CommonButton>編集する</CommonButton></router-link
  >
</template>
<script setup lang="ts">
import type { Liquor } from '@/graphQL/Liquor/liquor';
import CommonButton from '@/components/parts/common/CommonButton.vue';

interface Props {
  liquor: Liquor;
}

const { liquor } = defineProps<Props>();

console.log('liquor;', liquor);
</script>
<style scoped>
p.title {
  font-size: 150%;
  font-weight: bold;
}
</style>
