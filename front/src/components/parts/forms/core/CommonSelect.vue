<template>
  <label
    :for="props.name"
    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
    >{{ props.label }}</label
  >
  <select
    :id="props.name"
    :name="props.name"
    :disabled="props.isDisabled"
    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-1.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
  >
    <option
      v-for="option in props.options"
      :key="option.value"
      :value="option.value"
    >
      {{ option.label }}
    </option>
  </select>
</template>

<script setup lang="ts">
import type { Option } from '@/type/common/Option';

interface SelectProps {
  label?: string;
  name: string;
  options: Option[];
  isDisabled?: boolean;
}

const props = defineProps<SelectProps>();
</script>
<style scoped></style>
