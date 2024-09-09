<!--サイドバー-->
<template>
  <aside>
    カテゴリから検索:
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
import { computed, type ComputedRef, onMounted, ref } from 'vue';
import CategoryParent from '@/components/layouts/main/sideBar/CategoryParent.vue';
import { useSelectedCategoryStore } from '@/stores/sidebar';
import { getDisplayCategoryIds } from '@/funcs/component/laouts/main/sideBar/sideBarFunc';

const sidebarStore = useSelectedCategoryStore();

const { fetch } = useQuery<Categories>(GET_QUERY);

const categoryList = ref<Category[] | null>();

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const { categories: response } = await fetch();
  categoryList.value = response;
});

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
