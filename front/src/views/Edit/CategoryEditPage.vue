<template>
  <CategoryEdit v-if="!isLoading" :history-data="category" />
</template>

<script setup lang="ts">
import CategoryEdit from '@/components/templates/post/CategoryEdit.vue';
import useQuery from '@/funcs/composable/useQuery';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import {
  type CategoryHistoryData,
  GET_LOGS_FOR_ROLLBACK,
  type HistoryResponse,
} from '@/graphQL/Category/categoryLog';

const isLoading = ref<boolean>(true);
const category = ref<CategoryHistoryData | null>(null); //フィールドにあるカテゴリ情報

const route = useRoute(); // 現在のルートを取得
const { fetch } = useQuery<HistoryResponse>(GET_LOGS_FOR_ROLLBACK);

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const id: string | null = route.params.id as string; // ルートパラメータからidを取得
  if (!id) {
    isLoading.value = false;
  }
  await fetch({
    variables: {
      id: Number(id),
    },
    fetchPolicy: 'no-cache',
  })
    .then((response) => {
      category.value = response.histories;
    })
    .finally(() => {
      isLoading.value = false;
    });
});
</script>

<style scoped></style>
