<template>
  <CommonButton
    v-if="isReady"
    v-bind="props"
    :is-disabled="!isValid"
    :class="props.class"
    ><slot>送信</slot></CommonButton
  >
</template>

<script setup lang="ts">
import { useIsFormValid } from 'vee-validate';
import { onMounted, ref, watch } from 'vue';

import CommonButton from '@/components/parts/common/CommonButton.vue';
import type { ButtonProps } from '@/type/component/parts/ButtonProps';

const props = defineProps<ButtonProps>();

const isReady = ref(false); // フォームの検証が完了するまでボタンを表示しない
const isValid = ref<boolean>(false); //初期状態はfalseで定義する(直接useIsFormValidを使うと、非同期なので一瞬validになってしまう)
const formValid = useIsFormValid();

onMounted(() => {
  const valid = useIsFormValid();

  // フォームの検証結果が取得されたら isValid を更新し、isReady を true に設定
  isValid.value = valid.value;
  isReady.value = true; // 準備が整ったら描画
});

watch(formValid, (value) => {
  isValid.value = value;
});
</script>

<style scoped></style>
