<template>
  <CategoryEdit v-if="!isLoading" :history-data="category" />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

import useQuery from '@/funcs/composable/useQuery/useQuery';
import { isEmpty } from '@/funcs/util/isEmpty';
import {
  type CategoryHistoryData,
  GET_LOGS_FOR_ROLLBACK,
  type HistoryResponse,
} from '@/graphQL/Category/categoryLog';
import CategoryEdit from '@/views/Edit/CategoryEdit/form/CategoryEdit.vue';

const isLoading = ref<boolean>(true);
const category = ref<CategoryHistoryData | null>(null); //フィールドにあるカテゴリ情報

const route = useRoute(); // 現在のルートを取得
const { fetch } = useQuery<HistoryResponse>(GET_LOGS_FOR_ROLLBACK);

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const id: string = route.params.id as string; // ルートパラメータからidを取得
  if (isEmpty(id)) {
    isLoading.value = false;
    return;
  }
  await fetch(
    {
      id: Number(id),
    },
    {
      fetchPolicy: 'no-cache',
    },
  )
    .then((response) => {
      category.value = response.histories;
    })
    .finally(() => {
      isLoading.value = false;
    });
});
</script>

<style scoped></style>
