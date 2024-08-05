<!--汎用セレクトボックス-->
<template>
  <FormField v-bind="props" as="select" :multiple="props.isMultiple">
    <option v-if="isDisplayEmpty" value="" :disabled="props.isRequired">
      未選択
    </option>
    <option v-for="option in options" :key="option.value" :value="option.value">
      {{ option.label }}
    </option>
  </FormField>
</template>

<script setup lang="ts">
import FormField, {
  type FormFieldProps,
} from '@/components/parts/forms/core/FormField.vue';

interface SelectProps {
  isMultiple?: boolean;
  isRequired?: boolean;
  isDisplayEmpty?: boolean;
  options: Option[];
}

export interface Option {
  label: string;
  value: string | number;
}

export type SelectFieldProps = SelectProps & Omit<FormFieldProps, 'as'>;
const props = defineProps<SelectFieldProps>();
</script>

<style scoped></style>
