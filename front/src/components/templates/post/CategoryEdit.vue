<template>
  <VForm
    @submit="onSubmit"
    :v-if="initialValues"
    :validation-schema="validationSchema"
    :initial-values="initialValues"
  >
    <CategorySelect
      :name="FormKeys.PARENT"
      :initial-id="initialValues[FormKeys.PARENT]"
    />
    <ErrorMessage :name="FormKeys.PARENT" />
    <FormField :name="FormKeys.ID" type="hidden" />
    <FormField :name="FormKeys.NAME" label="名前" />
    <FormField :name="FormKeys.DESCRIPTION" label="説明" as="textarea" />
    <img
      v-if="initialData?.imageBase64"
      :src="`data:image/jpg;base64,${initialData.imageBase64}`"
      alt="画像"
    />
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
import FormField from '@/components/parts/forms/core/FormField.vue';
import CategorySelect from '@/components/blocks/common/forms/advance/CategorySelect.vue';
import SubmitButton from '@/components/parts/common/SubmitButton.vue';
import { useToast } from '@/funcs/composable/useToast';
import type { ToastCommand } from '@/plugins/toast';
import { useApiMutation } from '@/funcs/composable/useApiMutation';
import { useLoading } from 'vue-loading-overlay';
import { ref } from 'vue';
import type { CategoryForEdit } from '@/graphQL/Liquor/categories';
import CategoryPostAPIType, {
  type CategoryRequest,
  type CategoryResponse,
} from '@/type/api/APIType/post/CategoryForm';
import {
  FormKeys,
  type FormValues,
  generateInitialValues,
  validationSchema,
} from '@/forms/Post/CategoryForm';

// propsから受け取る初期値
const { initialData } = defineProps<{
  initialData: CategoryForEdit | null;
}>();

//必要な関数をインポート
const { mutateAsync } = useApiMutation<CategoryRequest, CategoryResponse>(
  CategoryPostAPIType,
);
const toast: ToastCommand = useToast();
const loading = useLoading();

// 初期値を定義
const initialValues = ref<FormValues>(generateInitialValues(initialData));

async function onSubmit(values: FormValues): Promise<void> {
  const loader = loading.show();
  //Categoryが空はバリデーションで弾かれる想定なのでキャスト
  await mutateAsync(<CategoryRequest>values, {
    onSuccess(value) {
      console.log('レスポンス：', value);
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
