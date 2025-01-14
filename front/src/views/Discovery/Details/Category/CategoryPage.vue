<template>
  <CategoryDetail v-if="category" :category="category" />
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';

import useQuery from '@/funcs/composable/useQuery/useQuery';
import {
  type Category,
  type CategoryResponse,
  GET_DETAIL,
} from '@/graphQL/Category/categories';
import { useSelectedCategoryStore } from '@/stores/sidebar';
import CategoryDetail from '@/views/Discovery/Details/Category/CategoryDetail.vue';

const route = useRoute(); // 現在のルートを取得
const sidebarStore = useSelectedCategoryStore();
const { fetch } = useQuery<CategoryResponse<Category>>(GET_DETAIL);

const category = ref<Category | null>(null);

const isNoCache: boolean = window.history.state?.noCache ?? false; //TODO:何故か常にtrueになってる...？

// データフェッチ
const fetchData = async (id: number): Promise<void> => {
  const { category: response } = await fetch(
    {
      id,
    },
    {
      fetchPolicy: isNoCache ? 'no-cache' : undefined, //更新直後だとキャッシュが残っているため、キャッシュを無効化
    },
  );
  category.value = response;
  sidebarStore.updateContent(response.id);
};

watch(
  () => route.params.id, // ルートのパスやクエリ、パラメータなどを監視
  (to) => {
    // ルートが変更された際に実行される処理
    const id = to as string; // ルートパラメータからidを取得
    if (!id) {
      return;
    }
    fetchData(Number(id));
  },
  { immediate: true }, // 初回レンダリング時に実行される
);
</script>

<style scoped></style>
