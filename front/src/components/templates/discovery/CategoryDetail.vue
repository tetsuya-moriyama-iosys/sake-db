<template>
  カテゴリー詳細
  <div>
    <p>{{ category.name }}</p>
  </div>
  <router-link :to="{ name: 'LiquorEdit', params: { id: category.id } }">
    <CommonButton>編集する</CommonButton></router-link
  >
</template>
<script setup lang="ts">
import CommonButton from '@/components/parts/common/CommonButton.vue';
import type { Category } from '@/graphQL/Liquor/categories';

interface Props {
  category: Category;
}
const { category } = defineProps<Props>();
</script>

<style scoped></style>
