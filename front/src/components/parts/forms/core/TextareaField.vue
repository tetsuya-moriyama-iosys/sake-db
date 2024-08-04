<!--汎用セレクトボックス-->
<template>
  <FormField v-bind="props" as="textarea" />
</template>

<script setup lang="ts">
import FormField, {
  type FormFieldProps,
} from '@/components/parts/forms/core/FormField.vue';

interface TextareaProps {}

export type SelectFieldProps = TextareaProps & Omit<FormFieldProps, 'as'>;
const props = defineProps<SelectFieldProps>();
</script>

<style scoped></style>
