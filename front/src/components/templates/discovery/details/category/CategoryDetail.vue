<template>
  カテゴリー詳細
  <div>
    <p>{{ category.name }}</p>
  </div>
  <img v-if="category.imageUrl" :src="category.imageUrl" alt="画像" />
  <div>
    <p>{{ category.description }}</p>
  </div>
  <CategoryChildren
    v-if="category.children"
    :category-list="category.children"
  />
  <router-link :to="{ name: 'CategoryEdit', params: { id: category.id } }">
    <CommonButton>編集する</CommonButton></router-link
  >
</template>
<script setup lang="ts">
import CommonButton from '@/components/parts/common/CommonButton.vue';
import type { Category } from '@/graphQL/Category/categories';
import CategoryChildren from '@/components/blocks/discovery/categoryDetail/CategoryChildren.vue';

interface Props {
  category: Category;
}
const { category } = defineProps<Props>();
</script>

<style scoped></style>
