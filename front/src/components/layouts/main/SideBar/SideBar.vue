<!--サイドバー-->
<template>
  <aside>
    カテゴリから検索
    <CategoryParent
      v-for="category in categoryList"
      :key="category.id"
      :category="category"
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
import { onMounted, ref } from 'vue';
import CategoryParent from '@/components/layouts/main/SideBar/CategoryParent.vue';

const { fetch } = useQuery<Categories>(GET_QUERY);

const categoryList = ref<Category[] | null>();

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const { categories: response } = await fetch();
  categoryList.value = response;
});
</script>

<style scoped>
aside {
  width: 180px;
  height: 100%;
  background-color: aquamarine;
}
</style>
