<template>
  <CommonButton v-bind="props" :is-disabled="!isValid"><slot /></CommonButton>
</template>

<script setup lang="ts">
import CommonButton from '@/components/parts/common/CommonButton.vue';
import type { ButtonProps } from '@/type/component/parts/ButtonProps';
import { useIsFormValid } from 'vee-validate';

const props = defineProps<ButtonProps>();

const isValid = useIsFormValid(); // 親コンポーネントで定義されたフォームの状態を取得
</script>

<style scoped></style>
