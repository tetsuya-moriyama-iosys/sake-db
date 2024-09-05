<template>
  <FromCategory
    :key="route.params.id as string"
    v-if="liquors"
    :liquors="liquors"
  />
</template>
<script setup lang="ts">
import { useRoute } from 'vue-router';
import useQuery from '@/funcs/composable/useQuery';
import {
  LIQUOR_LIST_FROM_CATEGORY,
  type ListResponse,
} from '@/graphQL/Liquor/liquor';
import type { Liquor } from '@/graphQL/Index/random';
import { onMounted, ref, watch } from 'vue';
import FromCategory from '@/components/templates/discovery/FromCategory.vue';

const route = useRoute(); // 現在のルートを取得
const { fetch } = useQuery<ListResponse>(LIQUOR_LIST_FROM_CATEGORY);

const liquors = ref<Liquor[] | null>(null);

// `watch` を使ってルートパラメータの変更を監視
watch(
  () => route.fullPath, // ルートのパスやクエリ、パラメータなどを監視
  (to) => {
    // ルートが変更された際に実行される処理
    const lastSegment = to.split('/').pop(); // 最後のセグメントを取得
    fetchData(Number(lastSegment));
  },
);

// 初期のデータフェッチ
const fetchData = async (id: number): Promise<void> => {
  const { listFromCategory: response } = await fetch({
    variables: {
      id,
    },
  });
  liquors.value = response;
};

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  void fetchData(Number(route.params.id as string));
});
</script>

<style scoped></style>
