<template>
  <p v-if="route.params.id">お酒ページ編集</p>
  <p v-else>お酒ページ作成</p>
  <div class="flex">
    <div class="flex-1">
      <LiquorForm
        :initial-data="initialValues"
        :version-no="historyData?.now.versionNo ?? null"
      />
    </div>
    <div>
      <LiquorLogs
        v-if="historyData?.histories"
        :logs="[historyData.now, ...historyData.histories]"
        @selectLog="reflectLog"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';

import type { Liquor } from '@/graphQL/Liquor/liquor';
import type { LiquorHistoryData } from '@/graphQL/Liquor/liquorLog';
import LiquorForm from '@/views/Edit/LiquorEdit/form/LiquorForm.vue';
import LiquorLogs from '@/views/Edit/LiquorEdit/form/LiquorLogs.vue';

// propsから受け取る初期値
const props = defineProps<{
  historyData: LiquorHistoryData | null;
}>();

const route = useRoute(); // 現在のルートを取得

// 初期値を定義
const initialValues = ref<Liquor | null>(props.historyData?.now ?? null);

//バージョン履歴をコピーする時に呼ばれる関数
const reflectLog = (log: Liquor) => {
  initialValues.value = { ...log, id: props.historyData!.now.id }; //過去のデータをそのまま初期値として代入する(IDだけは現在の値で上書き。ここが呼び出された次点で初期値は存在しないとおかしいのでアサーション)
};

watch(
  () => props.historyData,
  (liquor: LiquorHistoryData | null) => {
    initialValues.value = liquor?.now ?? null;
  },
);
</script>

<style scoped></style>
