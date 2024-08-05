<template>
  <div>
    <div v-if="label">
      {{ label }}
    </div>
    <div>
      <Field v-if="!Boolean($slots.default)" v-bind="props" />
      <!--セレクトボックス用(inputはslotがあると正常に描画されない)-->
      <Field v-if="$slots.default" v-bind="props">
        <slot></slot>
      </Field>
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
