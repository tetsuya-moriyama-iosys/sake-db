<template>
  <div class="ml-2" v-if="props.displayIds.includes(props.category.id)">
    ┗<span
      v-if="sidebarStore.content != props.category.id"
      class="category-name"
      ><router-link
        class="inline-block"
        :to="{ name: 'CategoryNarrowDown', params: { id: props.category.id } }"
        >{{ props.category.name }}</router-link
      ></span
    >
    <span v-else class="category-name selected font-bold">
      {{ props.category.name }}
    </span>
    <CategoryParent
      v-for="child in props.category.children"
      :key="child.id"
      :category="child"
      :display-ids="props.displayIds"
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

const props = defineProps<Props>();

const sidebarStore = useSelectedCategoryStore();
</script>

<style scoped></style>
