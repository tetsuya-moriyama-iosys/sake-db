<template>
  <VForm
    @submit="onSubmit"
    :v-if="initialValues"
    :validation-schema="validationSchema"
    :initial-values="initialValues"
  >
    <CategorySelect
      :name="FormKeys.CATEGORY"
      :initial-id="initialValues[FormKeys.CATEGORY]"
    />
    <ErrorMessage :name="FormKeys.CATEGORY" />
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
import {
  ErrorMessage,
  Form as VForm,
  type SubmissionHandler,
} from 'vee-validate';
import {
  FormKeys,
  type FormValues,
  generateInitialValues,
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
import type { LiquorForEdit } from '@/graphQL/Liquor/liquor';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import type { AxiosResponse } from 'axios';

// propsから受け取る初期値
const { initialData } = defineProps<{
  initialData: LiquorForEdit | null;
}>();

//必要な関数をインポート
const { mutateAsync } = useApiMutation<PostRequest, PostResponse>(PostAPIType);
const router = useRouter();
const toast: ToastCommand = useToast();
const loading = useLoading();

// 初期値を定義
const initialValues = ref<FormValues>(generateInitialValues(initialData));

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (
  values: FormValues,
): Promise<void> => {
  const loader = loading.show();
  //Categoryが空はバリデーションで弾かれる想定なのでキャスト
  await mutateAsync(<PostRequest>values, {
    onSuccess(value: AxiosResponse<PostResponse>) {
      toast.showToast({
        message: '登録が成功しました！',
      });
      router.push({
        name: 'LiquorDetail',
        params: { id: value.data.id },
        state: { noCache: true },
      });
    },
    onSettled() {
      loader.hide();
    },
  });
};
</script>

<style scoped></style>
