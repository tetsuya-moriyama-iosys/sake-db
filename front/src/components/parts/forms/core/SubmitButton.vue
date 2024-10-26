<template>
  <CommonButton
    v-bind="props"
    :is-disabled="!isReady || !isValid || !isDirty"
    :class="props.class"
    ><slot>送信</slot></CommonButton
  >
</template>

<script setup lang="ts">
import { useIsFormDirty, useIsFormValid } from 'vee-validate';
import { onMounted, ref, watch } from 'vue';

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import type { ButtonProps } from '@/components/parts/common/CommonButton/type';

const props = defineProps<ButtonProps>();

const isReady = ref(false); // フォームの検証が完了するまでボタンを表示しない
const isValid = ref<boolean>(false); //初期状態はfalseで定義する(直接useIsFormValidを使うと、非同期なので一瞬validになってしまう)
const formValid = useIsFormValid(); //memo:なんか初期値はtrueらしいので、描画直後はtrueになる。ボタンのチラつきを抑えるには、これを考慮したオプションを渡す必要がある
const isDirty = useIsFormDirty();

onMounted(() => {
  // フォームの検証結果が取得されたら isValid を更新し、isReady を true に設定
  isValid.value = formValid.value;
  isReady.value = true; // 準備が整ったら描画
});

watch(formValid, (value) => {
  isValid.value = value;
});

watch(
  isDirty,
  () => {
    console.log('isDirty.value:', isDirty.value);
  },
  {
    immediate: true,
  },
);
</script>

<style scoped></style>
