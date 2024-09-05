<!--カテゴリ一覧のパンくずリスト-->
<template>
  <div class="category-trail">
    <div
      v-for="(category, index) in categoryTrails"
      :key="category.id"
      class="inline-block category-trail-item"
    >
      <router-link
        :to="{ name: 'CategoryDetail', params: { id: category.id } }"
      >
        {{ category.name }}
      </router-link>
      <span v-if="index !== categoryTrails.length - 1" class="mx-2">
        &gt;
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  categoryTrails: {
    id: number;
    name: string;
  }[];
}
const { categoryTrails } = defineProps<Props>();
</script>

<style scoped></style>
