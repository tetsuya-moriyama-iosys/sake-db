<template>
  <div class="ml-2">
    ┗<router-link
      class="inline-block"
      :to="{ name: 'CategoryNarrowDown', params: { id: category.id } }"
      ><p>{{ category.name }}</p></router-link
    >
    <CategoryParent
      v-for="child in category.children"
      :key="child.id"
      :category="child"
    />
  </div>
</template>

<script setup lang="ts">
import type { Category } from '@/graphQL/Category/categories';

interface Props {
  category: Category;
}

const { category } = defineProps<Props>();
</script>

<style scoped></style>
