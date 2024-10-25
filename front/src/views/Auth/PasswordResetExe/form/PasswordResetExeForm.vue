<template>
  <Form @submit="onSubmit" :validation-schema="validationSchema">
    <FormField :name="PASSWORD_NAME" label="パスワード" type="password" />
    <FormField
      :name="PASSWORD2_NAME"
      label="パスワード(再入力)"
      type="password"
    />
    <SubmitButton />
  </Form>
</template>

<script setup lang="ts">
import { Form, type SubmissionHandler } from 'vee-validate';
import { useRouter } from 'vue-router';
import * as yup from 'yup';
import { object, string } from 'yup';

import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';
import { useMutation } from '@/funcs/composable/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import {
  PASSWORD_RESET_EXE,
  type ResetEmailExeResponse,
} from '@/graphQL/Auth/auth';
import { ToastType } from '@/plugins/toast';
import { useUserStore } from '@/stores/userStore';

const props = defineProps<{
  token: string;
}>();

//メールアドレス以外現状要らないので、定数もここで定義
const PASSWORD_NAME = 'password';
const PASSWORD2_NAME = 'password_2';

// yupのバリデーションスキーマを定義()
const validationSchema = object({
  [PASSWORD_NAME]: string()
    .required('パスワードは必須です')
    .min(6, 'パスワードは6文字以上である必要があります'),
  [PASSWORD2_NAME]: string()
    .required('確認用パスワードは必須です')
    .min(6, 'パスワードは6文字以上である必要があります')
    .oneOf([yup.ref(PASSWORD_NAME)], 'パスワードが一致しません'),
});

const toast = useToast();
const router = useRouter();
const userStore = useUserStore();

const { execute } = useMutation<ResetEmailExeResponse>(PASSWORD_RESET_EXE, {
  isUseSpinner: true,
});

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (values: {
  [PASSWORD_NAME]: string;
  [PASSWORD2_NAME]: string; //使わないけど一応
}) => {
  await execute({
    token: props.token,
    password: values[PASSWORD_NAME],
  }).then((response: ResetEmailExeResponse) => {
    toast.showToast({
      message: 'パスワードリセットが完了しました。',
      type: ToastType.Success,
    });
    userStore.setUserData(response.resetExe); //ストアの情報を更新する
    router.push({ name: 'Index' });
  });
};
</script>

<style scoped></style>
