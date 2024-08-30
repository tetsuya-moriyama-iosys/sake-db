<!--酒詳細ページ-->

<template>
  <p class="title">{{ liquor.name }}</p>
  <img v-if="liquor.imageUrl" :src="liquor.imageUrl" />
</template>
<script setup lang="ts">
import type { Liquor } from '@/graphQL/Liquor/liquor';

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
