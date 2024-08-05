<template>
  <div>新規投稿</div>
  <VForm @submit="onSubmit" :validation-schema="validationSchema">
    <CategorySelect
      :name="FormKeys.CATEGORY"
      :initial-id="values[FormKeys.CATEGORY]"
    />
    <FormField :name="FormKeys.TITLE" label="名前" />
    <FormField :name="FormKeys.DESCRIPTION" label="説明" />
    <SelectField
      name="test"
      label="説明2"
      :options="[
        {
          value: '1',
          label: 'test1',
        },
      ]"
    />
    <button type="submit">Submit</button>
  </VForm>
</template>

<script setup lang="ts">
import { Form as VForm, useForm } from 'vee-validate';
import {
  FormKeys,
  initialValues,
  validationSchema,
} from '@/forms/Post/CreatePostForm';
import FormField from '@/components/parts/forms/core/FormField.vue';
import CategorySelect from '@/components/blocks/common/forms/advance/CategorySelect.vue';
import SelectField from '@/components/parts/forms/core/SelectField.vue';

const { values } = useForm({
  initialValues: initialValues,
  validationSchema: validationSchema,
});

const onSubmit = () => {
  console.log(values);
};
</script>

<style scoped></style>
