<template>
  <VForm @submit="onSubmit" ref="form" :validation-schema="validationSchema">
    <CategorySelect
      label="カテゴリ"
      :name="FormKeys.CATEGORY"
      :initial-id="initialParentId"
    />
    <FormField :name="FormKeys.ID" type="hidden" />
    <FormField :name="FormKeys.VERSION_NO" type="hidden" />
    <FormField :name="FormKeys.SELECTED_VERSION_NO" type="hidden" />
    <FormField :name="FormKeys.NAME" label="名前" />
    <FormField :name="FormKeys.DESCRIPTION" label="説明" as="textarea" />
    <UploadWithImage
      v-slot="{ setImage }"
      :name="FormKeys.IMAGE"
      :default="initialData?.imageBase64"
    >
      <!-- trigger を ref に保存 -->
      <template v-if="setTrigger(setImage)"></template>
    </UploadWithImage>
    <SubmitButton>登録</SubmitButton>
  </VForm>
</template>

<script setup lang="ts">
import { Form as VForm, type SubmissionHandler } from 'vee-validate';
import {
  FormKeys,
  type FormValues,
  generateInitialValues,
  validationSchema,
} from '@/forms/define/post/CreatePostForm';
import FormField from '@/components/parts/forms/core/FormField.vue';
import CategorySelect from '@/components/parts/forms/common/CategorySelect.vue';
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
import { computed, type ComputedRef, onMounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import type { AxiosResponse } from 'axios';
import UploadWithImage from '@/components/parts/forms/common/UploadWithImage.vue';

// propsから受け取る初期値
const props = defineProps<{
  initialData: LiquorForEdit | null;
  versionNo: number | null;
}>();

//必要な関数をインポート
const { mutateAsync } = useApiMutation<PostRequest, PostResponse>(PostAPIType);
const router = useRouter();
const toast: ToastCommand = useToast();
const loading = useLoading();

const form = ref<InstanceType<typeof VForm> | null>(null); //Form内部に定義されているフォームメソッドにアクセスするのに必要

// trigger を保存する変数(↓typeなのにno-unused-varsが出るのでコメントアウト)
// eslint-disable-next-line no-unused-vars
type SetImage = (value: string | undefined) => void;
let setImage: SetImage | null = null;

// 子コンポーネントから受け取った trigger をセットする関数
const setTrigger = (fn: SetImage) => {
  setImage = fn;
  return true; // v-if のために true を返す
};

//初期データが変更されたら、フォームをリセットする
const resetForm = () => {
  if (setImage == null) {
    //通常ここには到達しないはず
    return;
  }
  form.value?.resetForm({
    values: {
      ...generateInitialValues(props.initialData),
      [FormKeys.VERSION_NO]: props.versionNo,
    },
  });
  setImage(props.initialData?.imageBase64);
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
  () => props.initialData?.categoryId ?? null,
);

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
