<!--リコメンドリストレイアウト-->
<template>
  <div v-if="dataList.length > 0">
    リコメンドリスト
    <div id="random-pickup-area">
      <CardContainer>
        <RecommendCard
          v-for="data in dataList"
          :data="data"
          :key="data.liquor.id"
        />
      </CardContainer>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

import CardContainer from '@/components/parts/common/CardContainer.vue';
import { DEFAULT_GET_LIMIT } from '@/const/indexConsts';
import useQuery from '@/funcs/composable/useQuery';
import {
  type Recommend,
  RECOMMEND_LIST_FROM_BOOKMARK,
  type RecommendLiquorList,
} from '@/graphQL/Index/recommends';
import RecommendCard from '@/views/Index/RecommendCard.vue';

const { fetch } = useQuery<RecommendLiquorList>(RECOMMEND_LIST_FROM_BOOKMARK, {
  isAuth: true,
});

const dataList = ref<Recommend[]>([]);

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const { getRecommendLiquorList: response } = await fetch({
    limit: DEFAULT_GET_LIMIT,
  });
  dataList.value = response;
});
</script>

<style scoped>
div#random-pickup-area {
  margin: auto;
  padding: 2em;
}
</style>
