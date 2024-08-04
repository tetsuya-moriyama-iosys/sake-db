<!--汎用テキストエリア-->
<template>
  <FormField v-bind="props" as="textarea" />
</template>

<script setup lang="ts">
import FormField, {
  type FormFieldProps,
} from '@/components/parts/forms/core/FormField.vue';

interface TextareaProps {}

export type TextareaFieldProps = TextareaProps & Omit<FormFieldProps, 'as'>;
const props = defineProps<TextareaFieldProps>();
</script>

<style scoped></style>
