<template>
  <div>
    <div v-if="label">
      {{ label }}
    </div>
    <div>
      <Field
        :key="props.name"
        v-if="!Boolean($slots.default)"
        v-bind="props"
        :placeholder="props.placeholder"
        v-slot="{ field }"
      >
        <component
          v-bind="field"
          :is="props.as ? props.as : 'input'"
          :type="props.type"
          :class="classes"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-1.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          :placeholder="props.placeholder"
          :disabled="props.readonly"
        />
      </Field>
      <!--セレクトボックス用(inputはslotがあると正常に描画されない)-->
      <Field v-if="$slots.default" v-bind="props"> <slot></slot> </Field>
    </div>
    <div class="error">
      <ErrorMessage :name="props.name" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ErrorMessage, Field } from 'vee-validate';

export interface FormFieldProps {
  label?: string;
  name: string;
  placeholder?: string;
  as?: 'input' | 'textarea' | 'select'; //Fieldのas属性
  classes?: string; //Fieldのclass属性
  readonly?: boolean;
  type?:
    | 'text'
    | 'hidden'
    | 'email'
    | 'password'
    | 'number'
    | 'url'
    | 'tel'
    | 'date'
    | 'time'
    | 'datetime-local'
    | 'month'
    | 'week'
    | 'search'
    | 'file'
    | 'color'; //inputのtype属性
}

const props = defineProps<FormFieldProps>();
</script>

<style scoped>
div.error {
  color: red;
  font-size: 75%;
}
</style>
