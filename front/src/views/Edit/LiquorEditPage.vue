<template>
  <LiquorEdit v-if="!isLoading" :history-data="liquor" />
</template>

<script setup lang="ts">
import useQuery from '@/funcs/composable/useQuery';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import LiquorEdit from '@/components/templates/post/LiquorEdit.vue';
import {
  GET_LOGS_FOR_ROLLBACK,
  type HistoryResponse,
  type LiquorHistoryData,
} from '@/graphQL/Liquor/liquorLog';
import { isEmpty } from '@/funcs/util/isEmpty';

const isLoading = ref<boolean>(true);
const liquor = ref<LiquorHistoryData | null>(null); //フィールドにあるカテゴリ情報

const route = useRoute(); // 現在のルートを取得
const { fetch } = useQuery<HistoryResponse>(GET_LOGS_FOR_ROLLBACK);

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const id: string = route.params.id as string; // ルートパラメータからidを取得
  if (isEmpty(id)) {
    isLoading.value = false;
    return;
  }
  await fetch({
    variables: {
      id,
    },
    fetchPolicy: 'no-cache',
  })
    .then((response) => {
      liquor.value = response.liquorHistories;
    })
    .finally(() => {
      isLoading.value = false;
    });
});
</script>

<style scoped></style>
