<template>
  <CommonButton
    v-bind="props"
    :is-disabled="!isValid || !isReady"
    :class="props.class"
    ><slot>送信</slot></CommonButton
  >
</template>

<script setup lang="ts">
import { useIsFormValid, useIsSubmitting } from 'vee-validate';
import { ref, watch } from 'vue';

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import type { ButtonProps } from '@/components/parts/common/CommonButton/type';

const props = defineProps<ButtonProps>();

const isReady = ref(false); // フォームの検証が完了するまでボタンを表示しない
const isValid = ref<boolean>(false); //初期状態はfalseで定義する(直接useIsFormValidを使うと、非同期なので一瞬validになってしまう)
const formValid = useIsFormValid(); //memo:なんか初期値はtrueらしいので、描画直後はtrueになる。ボタンのチラつきを抑えるには、これを考慮したオプションを渡す必要がある
const isSubmitting = useIsSubmitting(); // これがないと何故か↓のロジックがうまくいかないので追加

// バリデーション結果を監視
// フォームが初期状態で評価されたかどうかを確認するためのwatch
watch([formValid, isSubmitting], ([valid, submitting]) => {
  if (!submitting) {
    isValid.value = valid;
    isReady.value = true; // バリデーションが終了したらボタンを表示
  }
});
</script>

<style scoped></style>
