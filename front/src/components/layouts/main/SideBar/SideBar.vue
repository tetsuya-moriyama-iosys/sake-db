<!--サイドバー-->
<template>
  <div class="container flex flex-col">
    <aside class="flex-1">
      カテゴリから検索
      <CategoryParent
        v-for="category in categoryList"
        :key="category.id"
        :category="category"
        :display-ids="filteredCategoryIdList"
      />
    </aside>
    <aside class="new-post">
      <router-link :to="{ name: 'CategoryEdit' }"
        >+新規カテゴリ追加</router-link
      >
    </aside>
  </div>
</template>

<script setup lang="ts">
import { computed, type ComputedRef, onMounted, ref, watch } from 'vue';

import CategoryParent from '@/components/layouts/main/SideBar/CategoryParent.vue';
import { getDisplayCategoryIds } from '@/components/layouts/main/SideBar/func/sideBarFunc';
import useQuery from '@/funcs/composable/useQuery';
import {
  type Categories,
  type Category,
  GET_QUERY,
} from '@/graphQL/Category/categories';
import { useSelectedCategoryStore } from '@/stores/sidebar';

const sidebarStore = useSelectedCategoryStore();

const { fetch } = useQuery<Categories>(GET_QUERY);

const categoryList = ref<Category[] | null>();

async function fetchData() {
  const { categories: response } = await fetch(null, {
    fetchPolicy: 'no-cache',
  });
  categoryList.value = [...response];
  sidebarStore.setReloadFlgFalse();
}

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  void fetchData();
});

watch(
  () => sidebarStore.isReloadFlg,
  () => {
    if (sidebarStore.isReloadFlg) {
      void fetchData();
    }
  },
);

// sidebarStore.contentに基づいてカテゴリをフィルタリングする
const filteredCategoryIdList: ComputedRef<number[]> = computed(() => {
  if (!categoryList.value) return []; //そもそも存在していなければ処理終了
  if (!sidebarStore.content) return categoryList.value.map((c) => c.id); // contentがない場合は全ての大カテゴリを返す
  return getDisplayCategoryIds(categoryList.value, sidebarStore.content);
});
</script>

<style scoped>
div.container {
  width: 180px;
  background-color: aquamarine;
}
</style>
