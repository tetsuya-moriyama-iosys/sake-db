<template>
  <p>カテゴリー作成</p>
  <CategoryEdit v-if="!isLoading" :initial-data="category" />
</template>

<script setup lang="ts">
import CategoryEdit from '@/components/templates/post/CategoryEdit.vue';
import useQuery from '@/funcs/composable/useQuery';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import {
  type CategoryForEdit,
  type CategoryResponse,
  GET_DETAIL_FOR_EDIT,
} from '@/graphQL/Category/categories';

const isLoading = ref<boolean>(true);
const category = ref<CategoryForEdit | null>(null);

const route = useRoute(); // 現在のルートを取得
const { fetch } =
  useQuery<CategoryResponse<CategoryForEdit>>(GET_DETAIL_FOR_EDIT);

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const id: string = route.params.id as string; // ルートパラメータからidを取得
  if (!id) {
    isLoading.value = false;
    return;
  }
  await fetch({
    variables: {
      id,
    },
  })
    .then((response) => {
      category.value = response.category;
    })
    .finally(() => {
      isLoading.value = false;
    });
});
</script>

<style scoped></style>
