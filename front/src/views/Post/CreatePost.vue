<template>
  <div>新規投稿</div>
  <VForm @submit="onSubmit" :validation-schema="validationSchema">
    <MasterCategorySelect
      :name="FormKeys.CATEGORY"
      :is-display-empty="true"
      :is-required="true"
    />
    <FormField :name="FormKeys.TITLE" label="名前" />
    <FormField :name="FormKeys.TITLE" label="説明" />
    <CategorySelect name="category2" />
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
import MasterCategorySelect from '@/components/blocks/common/forms/MasterCategorySelect.vue';
import FormField from '@/components/parts/forms/core/FormField.vue';
import CategorySelect from '@/components/blocks/common/forms/advance/CategorySelect.vue';

useForm({
  initialValues: initialValues,
  validationSchema: validationSchema,
});

const onSubmit = () => {
  //console.log(values);
  // resetForm();
};
</script>

<style scoped></style>
