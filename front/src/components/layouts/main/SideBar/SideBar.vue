<!--サイドバー-->
<template>
  <aside>
    カテゴリから検索
    <CategoryParent
      v-for="category in categoryList"
      :key="category.id"
      :category="category"
      :display-ids="filteredCategoryIdList"
    />
  </aside>
</template>

<script setup lang="ts">
import useQuery from '@/funcs/composable/useQuery';
import {
  type Categories,
  type Category,
  GET_QUERY,
} from '@/graphQL/Category/categories';
import { computed, type ComputedRef, onMounted, ref, watch } from 'vue';
import { useSelectedCategoryStore } from '@/stores/sidebar';
import { getDisplayCategoryIds } from '@/funcs/component/laouts/main/sideBar/sideBarFunc';
import CategoryParent from '@/components/layouts/main/sideBar/CategoryParent.vue';

const sidebarStore = useSelectedCategoryStore();

const { fetch } = useQuery<Categories>(GET_QUERY);

const categoryList = ref<Category[] | null>();

async function fetchData() {
  const { categories: response } = await fetch({
    fetchPolicy: 'no-cache',
  });
  categoryList.value = [...response];
}

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  void fetchData();
  sidebarStore.setReloadFlgFalse();
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
aside {
  width: 180px;
  height: 100%;
  background-color: aquamarine;
}
</style>
