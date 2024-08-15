<template>
  <div>新規投稿</div>
  <VForm @submit="onSubmit" :validation-schema="validationSchema">
    <CategorySelect
      :name="FormKeys.CATEGORY"
      :initial-id="formValues[FormKeys.CATEGORY]"
    />
    <FormField :name="FormKeys.TITLE" label="名前" />
    <FormField :name="FormKeys.DESCRIPTION" label="説明" />
    <button>Submit</button>
  </VForm>
</template>

<script setup lang="ts">
import { Form as VForm, useForm } from 'vee-validate';
import {
  FormKeys,
  type FormValues,
  initialValues,
  validationSchema,
} from '@/forms/Post/CreatePostForm';
import FormField from '@/components/parts/forms/core/FormField.vue';
import CategorySelect from '@/components/blocks/common/forms/advance/CategorySelect.vue';
import { useMutation } from '@vue/apollo-composable';
import { CREATE_POST_MUTATION } from '@/graphQL/Discovery/Post/query';

// Apollo ClientのuseMutationフックを使ってミューテーションを実行
const {
  mutate: createLiquor,
  onDone,
  onError,
} = useMutation(CREATE_POST_MUTATION);

const { values: formValues } = useForm<FormValues>({
  initialValues: initialValues,
  validationSchema: validationSchema,
});

async function onSubmit(values: FormValues): Promise<void> {
  console.log('values:', values);
  await createLiquor({
    name: values[FormKeys.TITLE],
    category_id: Number(values[FormKeys.CATEGORY]),
    description: values[FormKeys.DESCRIPTION],
  });

  onDone((response) => {
    console.log('Post created successfully:', response.data);
  });

  onError((error) => {
    console.error('Error creating post:', error);
  });
  //const response = useQuery();
}
</script>

<style scoped></style>
