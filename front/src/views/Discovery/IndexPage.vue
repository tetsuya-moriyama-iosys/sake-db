<!--トップページ-->
<template>IndexPage</template>

<script setup lang="ts">
import useQuery from '@/funcs/composable/useQuery';
import {
  RANDOM_RECOMMEND_LIST,
  type RecommendLiquorResponse,
} from '@/graphQL/Index/random';
import { onMounted } from 'vue';
import { DEFAULT_GET_LIMIT } from '@/const/indexConsts';

const { fetch } = useQuery<RecommendLiquorResponse>(RANDOM_RECOMMEND_LIST);

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const response: RecommendLiquorResponse = await fetch({
    variables: {
      limit: DEFAULT_GET_LIMIT,
    },
  });
  console.log(response);
});
</script>

<style scoped></style>
