<template>
  <div class="card-container">
    <img
      v-if="item.imageURL"
      :src="item.imageURL"
      class="m-auto"
      alt="商品画像"
    />
    <div class="m-auto">
      <a :href="item.URL" target="_blank">{{ displayName }}</a>
    </div>
    <div v-if="item.price">￥{{ item.price }}</div>
  </div>
</template>
<script setup lang="ts">
import type { AffiliateItem } from '@/graphQL/Amazon/affiliate';
import truncateString from '@/funcs/util/transform/truncateString';

interface Props {
  item: AffiliateItem; //商品名
}
const { item } = defineProps<Props>();

const displayName = truncateString({
  str: item.name,
  maxLength: 40,
});
</script>

<style scoped>
div.card-container {
  display: grid;
  grid-template-rows: 1fr 80px; /* 上側を自動 (1fr)、下側を固定長 (100px) */
  height: 100%; /* コンテナ全体の高さを指定（親要素の高さに依存） */

  border: 1px solid #777;
}
</style>
