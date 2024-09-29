<template>
  <LiquorEdit v-if="!isLoading" :history-data="liquor" />
</template>

<script setup lang="ts">
import useQuery from '@/funcs/composable/useQuery';
import { onMounted, ref, watch } from 'vue';
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
  await fetch(
    {
      id,
    },
    {
      fetchPolicy: 'no-cache',
    },
  )
    .then((response) => {
      liquor.value = response.liquorHistories;
    })
    .finally(() => {
      isLoading.value = false;
    });
});

//編集→新規投稿の導線があるが、考慮されていないのでパス監視を追加する
watch(
  () => route.params.id, // ルートのパスやクエリ、パラメータなどを監視
  (to) => {
    // ルートが変更された際に実行される処理
    const id = to as string; // ルートパラメータからidを取得
    if (!id) {
      liquor.value = null;
      return;
    }
  },
  { immediate: true }, // 初回レンダリング時に実行される
);
</script>

<style scoped></style>
