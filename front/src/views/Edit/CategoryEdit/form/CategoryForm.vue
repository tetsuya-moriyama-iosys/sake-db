<template>
  <VForm @submit="onSubmit" ref="form" :validation-schema="validationSchema">
    親カテゴリ
    <CategorySelect
      :name="FormKeys.PARENT"
      :initial-id="initialParentId"
      :readonly="readonly"
    />
    <FormField :name="FormKeys.ID" type="hidden" />
    <FormField :name="FormKeys.VERSION_NO" type="hidden" />
    <FormField :name="FormKeys.SELECTED_VERSION_NO" type="hidden" />
    <FormField :name="FormKeys.NAME" label="名前" :readonly="readonly" />
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
import type { AxiosResponse } from 'axios';
import { Form as VForm, type SubmissionHandler } from 'vee-validate';
import { computed, type ComputedRef, onMounted, ref, watch } from 'vue';
import { useLoading } from 'vue-loading-overlay';
import { useRouter } from 'vue-router';

import CategorySelect from '@/components/parts/forms/common/CategorySelect.vue';
import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';
import { useApiMutation } from '@/funcs/composable/useApiMutation';
import { useToast } from '@/funcs/composable/useToast';
import type { Category } from '@/graphQL/Category/categories';
import type { ToastCommand } from '@/plugins/toast';
import { useSelectedCategoryStore } from '@/stores/sidebar';
import CategoryPostAPIType, {
  type CategoryRequest,
  type CategoryResponse,
} from '@/type/api/APIType/post/CategoryForm';
import {
  FormKeys,
  type FormValues,
  generateInitialValues,
  validationSchema,
} from '@/views/Edit/CategoryEdit/form/CategoryForm';

// propsから受け取る初期値
const props = defineProps<{
  initialData: Category | null;
  versionNo: number | null;
  readonly: boolean;
}>();

//必要な関数をインポート
const { mutateAsync } = useApiMutation<CategoryRequest, CategoryResponse>(
  CategoryPostAPIType,
);
const router = useRouter();
const toast: ToastCommand = useToast();
const loading = useLoading();

const sideBarStore = useSelectedCategoryStore();

const form = ref<InstanceType<typeof VForm> | null>(null); //Form内部に定義されているフォームメソッドにアクセスするのに必要

//初期データが変更されたら、フォームをリセットする
const resetForm = () => {
  form.value?.resetForm({
    values: {
      ...generateInitialValues(props.initialData),
      [FormKeys.VERSION_NO]: props.versionNo,
    },
  });
};

//初期値が変更されたらフォームをリセットする(Formコンポーネントに依存しているので、初回はonMounted)
watch(
  () => props.initialData,
  () => {
    resetForm();
  },
);
onMounted(() => {
  resetForm();
});

const initialParentId: ComputedRef<number | null> = computed(
  () => props.initialData?.parent ?? null,
);

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (
  values: FormValues,
): Promise<void> => {
  const loader = loading.show();
  //Categoryが空はバリデーションで弾かれる想定なのでキャスト
  await mutateAsync(<CategoryRequest>values, {
    onSuccess(value: AxiosResponse<CategoryResponse>) {
      toast.showToast({
        message: '登録が成功しました！',
      });
      sideBarStore.reload();
      router.push({
        name: 'CategoryDetail',
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
