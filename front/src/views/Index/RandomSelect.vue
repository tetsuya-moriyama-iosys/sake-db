<!--ランダムセレクトリストレイアウト-->
<template>
  <div id="random-pickup-area">
    <CardContainer>
      <LiquorCard v-for="liquor in liquors" :liquor="liquor" :key="liquor.id" />
    </CardContainer>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

import LiquorCard from '@/components/blocks/cards/LiquorCard.vue';
import CardContainer from '@/components/parts/common/CardContainer.vue';
import { DEFAULT_GET_LIMIT } from '@/const/indexConsts';
import useQuery from '@/funcs/composable/useQuery';
import {
  type Liquor,
  RANDOM_RECOMMEND_LIST,
  type RecommendLiquorResponse,
} from '@/graphQL/Index/random';

const { fetch } = useQuery<RecommendLiquorResponse>(RANDOM_RECOMMEND_LIST);

const liquors = ref<Liquor[]>([]);

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const { randomRecommendList: response } = await fetch({
    limit: DEFAULT_GET_LIMIT,
  });
  liquors.value = response;
});
</script>

<style scoped>
div#random-pickup-area {
  margin: auto;
  padding: 2em;
}
</style>
