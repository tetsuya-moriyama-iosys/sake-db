<template>
  <p v-if="route.params.id">カテゴリー編集</p>
  <p v-else>カテゴリー作成</p>
  <div class="flex">
    <div class="flex-1">
      <CategoryForm
        :initial-data="initialValues"
        :version-no="historyData?.now.versionNo ?? null"
      />
    </div>
    <div>
      <CategoryLogs
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

import type { Category } from '@/graphQL/Category/categories';
import type { CategoryHistoryData } from '@/graphQL/Category/categoryLog';
import CategoryForm from '@/views/Edit/CategoryEdit/form/CategoryForm.vue';
import CategoryLogs from '@/views/Edit/CategoryEdit/form/CategoryLogs.vue';

// propsから受け取る初期値
const { historyData } = defineProps<{
  historyData: CategoryHistoryData | null;
}>();

const route = useRoute(); // 現在のルートを取得

// 初期値を定義
const initialValues = ref<Category | null>(historyData?.now ?? null);

const reflectLog = (log: Category) => {
  initialValues.value = { ...log }; //過去のデータをそのまま初期値として代入する
};
</script>

<style scoped></style>
