<template>
  カテゴリで検索：
  <span v-if="Number(route.params.id) < 10"> {{ categoryName }}</span>
  <router-link
    v-else
    :to="{ name: 'CategoryEdit', params: { id: route.params.id } }"
  >
    {{ categoryName }}</router-link
  >
  {{ categoryDescription }}
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
  type ListFromCategoryResponse,
} from '@/graphQL/Liquor/liquor';
import type { Liquor } from '@/graphQL/Index/random';
import { ref, watch } from 'vue';
import { useSelectedCategoryStore } from '@/stores/sidebar';
import FromCategory from '@/components/templates/discovery/FromCategory.vue';

const route = useRoute(); // 現在のルートを取得
const sidebarStore = useSelectedCategoryStore();
const { fetch } = useQuery<ListFromCategoryResponse>(LIQUOR_LIST_FROM_CATEGORY);

const liquors = ref<Liquor[] | null>(null);
const categoryName = ref<string>('');
const categoryDescription = ref<string>('');

// データフェッチ
const fetchData = async (id: number): Promise<void> => {
  sidebarStore.updateContent(id);
  const { listFromCategory: response } = await fetch({
    id,
  });
  liquors.value = response.liquors;
  categoryName.value = response.categoryName;
  categoryDescription.value = response.categoryDescription;
};

// `watch` を使ってルートパラメータの変更を監視
watch(
  () => route.params.id, // ルートのパスやクエリ、パラメータなどを監視
  (to) => {
    // ルートが変更された際に実行される処理
    const id = to as string; // ルートパラメータからidを取得
    if (!id) {
      return;
    }
    fetchData(Number(to));
  },
  { immediate: true }, // 初回レンダリング時に実行される
);
</script>

<style scoped></style>
