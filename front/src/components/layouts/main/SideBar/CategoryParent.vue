<template>
  <div class="ml-2" v-if="displayIds.includes(category.id)">
    ┗<span v-if="sidebarStore.content != category.id" class="category-name"
      ><router-link
        class="inline-block"
        :to="{ name: 'CategoryNarrowDown', params: { id: category.id } }"
        >{{ category.name }}</router-link
      ></span
    >
    <span v-else class="category-name selected font-bold">
      {{ category.name }}
    </span>
    <CategoryParent
      v-for="child in category.children"
      :key="child.id"
      :category="child"
      :display-ids="displayIds"
    />
  </div>
</template>

<script setup lang="ts">
import type { Category } from '@/graphQL/Category/categories';
import { useSelectedCategoryStore } from '@/stores/sidebar';

interface Props {
  category: Category;
  displayIds: number[];
}

const { category, displayIds } = defineProps<Props>();

const sidebarStore = useSelectedCategoryStore();
</script>

<style scoped></style>
