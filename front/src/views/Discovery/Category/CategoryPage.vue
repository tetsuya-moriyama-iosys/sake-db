<template>
  <CategoryDetail v-if="category" :category="category" />
</template>
<script setup lang="ts">
import { useRoute } from 'vue-router';
import useQuery from '@/funcs/composable/useQuery';
import { type DetailResponse, GET_DETAIL } from '@/graphQL/Liquor/categories';
import type { Category } from '@/type/common/liquor/Category';
import { onMounted, ref } from 'vue';
import CategoryDetail from '@/components/templates/discovery/CategoryDetail.vue';

const route = useRoute(); // 現在のルートを取得
const { fetch } = useQuery<DetailResponse>(GET_DETAIL);

const category = ref<Category | null>(null);

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const { category: response } = await fetch({
    variables: {
      id: route.params.id, // ルートパラメータからidを取得,
    },
  });

  category.value = response;
  console.log(response);
});
</script>

<style scoped></style>
