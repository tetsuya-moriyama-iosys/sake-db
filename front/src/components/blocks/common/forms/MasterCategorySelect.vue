<!--親カテゴリ選択セレクトボックス-->

<template>
  <SelectField v-bind="propsObj" />
</template>

<script setup lang="ts">
import SelectField, {
  type SelectFieldProps,
} from '@/components/parts/forms/core/SelectField.vue';
import {
  MasterCategory,
  MasterCategoryLabel,
} from '@/const/db/category/master';

const { label = 'メインカテゴリ', ...props } =
  defineProps<Omit<SelectFieldProps, 'options'>>();

//選択肢
const options = [
  {
    label: MasterCategoryLabel[MasterCategory.SPIRITS],
    value: MasterCategory.SPIRITS,
  },
  {
    label: MasterCategoryLabel[MasterCategory.BREW],
    value: MasterCategory.BREW,
  },
  {
    label: MasterCategoryLabel[MasterCategory.LIQUEUR],
    value: MasterCategory.LIQUEUR,
  },
  {
    label: MasterCategoryLabel[MasterCategory.OTHER],
    value: MasterCategory.OTHER,
  },
];

//オブジェクトを作らないとバインディングされないので詰め替える
const propsObj = { ...props, label, options };
</script>

<style scoped></style>
