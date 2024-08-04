<template>
  <div>
    <div v-if="label">
      {{ label }}
    </div>
    <div>
      <Field v-bind="props"><slot></slot></Field>
    </div>
    <div class="error">
      <ErrorMessage :name="props.name" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { Field, ErrorMessage } from 'vee-validate';

export interface FormFieldProps {
  label?: string;
  name: string;
  as?: 'input' | 'textarea' | 'select'; //Fieldのas属性
  type?:
    | 'text'
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
