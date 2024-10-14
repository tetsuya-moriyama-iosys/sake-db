<template>
  <div>
    <div v-for="log in logs" :key="log.versionNo" @click="handleClick(log)">
      バージョン：{{ log.versionNo }}
      {{ log.updatedAt ? format(log.updatedAt, 'yyyy/MM/dd HH:mm:ss') : '' }}
    </div>
  </div>
</template>

<script setup lang="ts">
// propsから受け取る初期値
import { format } from 'date-fns';

import type { Category } from '@/graphQL/Category/categories';

const { logs } = defineProps<{
  logs: Category[];
}>();

const emit = defineEmits(['selectLog']); // 親に送るイベントを定義

const handleClick = (log: Category) => {
  emit('selectLog', log); // 第二引数としてデータを渡す
};
</script>

<style scoped></style>
