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
    <div>
      <iframe
        v-if="liquor.youtube"
        width="560"
        height="315"
        :src="embedUrl ?? undefined"
        title="YouTube video player"
        allow="accelerometer; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
        allowfullscreen
      ></iframe>
    </div>
    <router-link :to="{ name: 'LiquorEdit', params: { id: liquor.id } }">
      <CommonButton>編集する</CommonButton></router-link
    >
    <FlavorMap :liquor="liquor" />
    <LiquorTags :liquor-id="liquor.id" />
    <AffiliateContainer :name="liquor.name" />
    <LiquorBoard :liquorId="liquor.id" />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

import AffiliateContainer from '@/components/blocks/common/amazon/AffiliateContainer.vue';
import FlavorMap from '@/components/blocks/FlavorMap/FlavorMap.vue';
import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import type { Liquor } from '@/graphQL/Liquor/liquor';
import LiquorBoard from '@/views/Discovery/Details/Liquor/board/LiquorBoard.vue';
import CategoryTrail from '@/views/Discovery/Details/Liquor/CategoryTrail.vue';
import LiquorTags from '@/views/Discovery/Details/Liquor/tag/LiquorTagArea.vue';

interface Props {
  liquor: Liquor;
}

const { liquor } = defineProps<Props>();

// YouTubeのURLをembed形式に変換するcomputedプロパティ
const embedUrl = computed<string | null>(() => {
  if (!liquor.youtube) return null;
  const videoIdMatch = liquor.youtube.match(
    /(?:youtube\.com\/watch\?v=|youtu\.be\/)([a-zA-Z0-9_-]{11})/,
  );
  return videoIdMatch
    ? `https://www.youtube.com/embed/${videoIdMatch[1]}`
    : null;
});
</script>

<style scoped>
p.title {
  font-size: 150%;
  font-weight: bold;
}

img.image {
  max-height: 300px;
  max-width: 500px;
}
</style>
