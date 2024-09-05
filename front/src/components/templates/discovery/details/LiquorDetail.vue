<!--酒詳細ページ-->

<template>
  <div v-if="liquor">
    <p class="title">{{ liquor.name }}</p>
    <CategoryTrail :category-trails="liquor.categoryTrail" />
    <img
      v-if="liquor.imageUrl"
      :src="liquor.imageUrl"
      class="image"
      alt="画像"
    />
    <div>
      {{ liquor.description }}
    </div>
    <router-link :to="{ name: 'LiquorEdit', params: { id: liquor.id } }">
      <CommonButton>編集する</CommonButton></router-link
    >
  </div>
</template>
<script setup lang="ts">
import type { Liquor } from '@/graphQL/Liquor/liquor';
import CommonButton from '@/components/parts/common/CommonButton.vue';
import CategoryTrail from '@/components/blocks/discovery/CategoryTrail.vue';

interface Props {
  liquor: Liquor;
}

const { liquor } = defineProps<Props>();
</script>
<style scoped>
p.title {
  font-size: 150%;
  font-weight: bold;
}

img.image {
  max-width: 500px;
}
</style>
