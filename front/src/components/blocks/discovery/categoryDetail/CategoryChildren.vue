<!--子カテゴリのリスト-->
<template>
  子カテゴリ一覧
  <div v-for="category in categoryList" :key="category.id">
    <router-link :to="{ name: 'CategoryDetail', params: { id: category.id } }"
      ><p>{{ category.name }}</p></router-link
    >
  </div>
</template>
<script setup lang="ts">
import type { Category } from '@/graphQL/Category/categories';

interface Props {
  categoryList: Category[];
}

const { categoryList } = defineProps<Props>();
</script>

<style scoped></style>
