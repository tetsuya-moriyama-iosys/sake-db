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
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import LiquorForm from '@/components/blocks/post/liquor/LiquorForm.vue';
import LiquorLogs from '@/components/blocks/post/liquor/LiquorLogs.vue';
import type { LiquorHistoryData } from '@/graphQL/Liquor/liquorLog';
import type { Liquor } from '@/graphQL/Liquor/liquor';

// propsから受け取る初期値
const { historyData } = defineProps<{
  historyData: LiquorHistoryData | null;
}>();

const route = useRoute(); // 現在のルートを取得

// 初期値を定義
const initialValues = ref<Liquor | null>(historyData?.now ?? null);

const reflectLog = (log: Liquor) => {
  initialValues.value = { ...log }; //過去のデータをそのまま初期値として代入する
};
</script>

<style scoped></style>
