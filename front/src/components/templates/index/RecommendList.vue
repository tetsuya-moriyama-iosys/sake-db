<!--リコメンドリストレイアウト-->
<template>
  <div id="random-pickup-area">
    <CardContainer>
      <RecommendCard
        v-for="data in dataList"
        :data="data"
        :key="data.liquor.id"
      />
    </CardContainer>
  </div>
</template>

<script setup lang="ts">
import useQuery from '@/funcs/composable/useQuery';
import { onMounted, ref } from 'vue';
import { DEFAULT_GET_LIMIT } from '@/const/indexConsts';
import CardContainer from '@/components/parts/common/CardContainer.vue';
import {
  type Recommend,
  RECOMMEND_LIST_FROM_BOOKMARK,
  type RecommendLiquorList,
} from '@/graphQL/Index/recommends';
import RecommendCard from '@/components/blocks/index/RecommendCard.vue';

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
