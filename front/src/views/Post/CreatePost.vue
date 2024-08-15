<template>
  <div>新規投稿</div>
  <VForm
    @submit="onSubmit"
    :validation-schema="validationSchema"
    :initial-values="initialValues"
  >
    <CategorySelect
      :name="FormKeys.CATEGORY"
      :initial-id="initialValues[FormKeys.CATEGORY]"
    />
    <FormField :name="FormKeys.TITLE" label="名前" />
    <FormField :name="FormKeys.DESCRIPTION" label="説明" />
    <FormField
      :name="FormKeys.IMAGE"
      type="file"
      as="input"
      label="画像"
      rules="required|image|size:5000"
    />
    <CommonSelect name="a" :options="[{ value: 'a', label: 'bb' }]" />
    <CommonButton :color="ColorType.Default">登録</CommonButton>
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
import CommonButton from '@/components/parts/common/CommonButton.vue';
import { ColorType } from '@/type/common/ColorType';
import CommonSelect from '@/components/parts/forms/core/CommonSelect.vue';

// Apollo ClientのuseMutationフックを使ってミューテーションを実行
const {
  mutate: createLiquor,
  onDone,
  onError,
} = useMutation(CREATE_POST_MUTATION);

useForm<FormValues>({
  validationSchema: validationSchema,
});

async function onSubmit(values: FormValues): Promise<void> {
  try {
    await createLiquor({
      name: values[FormKeys.TITLE],
      category_id: Number(values[FormKeys.CATEGORY]),
      description: values[FormKeys.DESCRIPTION],
    });
  } catch (e) {
    console.error(e);
  }

  onDone((response) => {
    console.log('Post created successfully:', response.data);
  });

  onError((error) => {
    console.error('Error creating post:', error);
  });
}
</script>

<style scoped></style>
