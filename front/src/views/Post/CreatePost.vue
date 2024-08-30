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
    <ErrorMessage :name="FormKeys.CATEGORY" />
    <FormField :name="FormKeys.TITLE" label="名前" />
    <FormField :name="FormKeys.DESCRIPTION" label="説明" />
    <FormField
      :name="FormKeys.IMAGE"
      type="file"
      as="input"
      label="画像"
      rules="required|image|size:5000"
    />
    <SubmitButton>登録</SubmitButton>
  </VForm>
</template>

<script setup lang="ts">
import { ErrorMessage, Form as VForm } from 'vee-validate';
import {
  FormKeys,
  type FormValues,
  initialValues,
  validationSchema,
} from '@/forms/Post/CreatePostForm';
import FormField from '@/components/parts/forms/core/FormField.vue';
import CategorySelect from '@/components/blocks/common/forms/advance/CategorySelect.vue';
import SubmitButton from '@/components/parts/common/SubmitButton.vue';
import { useToast } from '@/funcs/composable/useToast';
import type { ToastCommand } from '@/plugins/toast';
import { useApiMutation } from '@/funcs/composable/useApiMutation';
import PostAPIType, {
  type PostRequest,
  type PostResponse,
} from '@/type/api/APIType/post/PostForm';
import { useLoading } from 'vue-loading-overlay';

const { mutateAsync } = useApiMutation<PostRequest, PostResponse>(PostAPIType);
const toast: ToastCommand = useToast();
const loading = useLoading();

async function onSubmit(values: FormValues): Promise<void> {
  const loader = loading.show();
  await mutateAsync(values, {
    onSuccess() {
      toast.showToast({
        message: '登録が成功しました！',
      });
    },
    onSettled() {
      loader.hide();
    },
  });
}
</script>

<style scoped></style>
