<template>
  <Form
    @submit="onSubmit"
    :initial-values="initialValues"
    :validation-schema="validationSchema"
  >
    <FormField :name="FormKeys.MAIL" label="メールアドレス" type="email" />
    <FormField :name="FormKeys.PASSWORD" label="パスワード" type="password" />
    <SubmitButton>ログイン</SubmitButton>
  </Form>
  <router-link :to="{ name: 'PasswordReset' }">パスワードリセット</router-link>
</template>

<script setup lang="ts">
import { Form, type SubmissionHandler } from 'vee-validate';
import { useRouter } from 'vue-router';

import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';
import { useMutation } from '@/funcs/composable/useQuery/useQuery';
import { LOGIN } from '@/graphQL/Auth/auth';
import type { login, LoginMutation, Mutation } from '@/graphQL/auto-generated';
import { getAuthPayloadForUI } from '@/stores/userStore/type';
import { useUserStore } from '@/stores/userStore/userStore';
import {
  FormKeys,
  type FormValues,
  initialValues,
  validationSchema,
} from '@/views/Auth/Login/form/LoginForm';

const router = useRouter();
const userStore = useUserStore();
const { execute } = useMutation<LoginMutation>(LOGIN);

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (values: FormValues) => {
  await execute({
    input: {
      email: values[FormKeys.MAIL],
      password: values[FormKeys.PASSWORD],
    },
  }).then((res) => {
    //トークンをセットし、トップへリンク
    userStore.setUserData(getAuthPayloadForUI(res.login)); //ストアの情報を更新する
    router.push({ name: 'Index' });
  });
};
</script>

<style scoped></style>
