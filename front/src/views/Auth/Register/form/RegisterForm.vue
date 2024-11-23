<template>
  <Form
    @submit="onSubmit"
    :initial-values="initialValues"
    :validation-schema="validationSchema"
  >
    <FormField :name="FormKeys.NAME" label="名前" />
    <FormField :name="FormKeys.MAIL" label="メールアドレス" type="email" />
    <FormField :name="FormKeys.PASSWORD" label="パスワード" type="password" />
    <SubmitButton>登録</SubmitButton>
  </Form>
</template>

<script setup lang="ts">
import { Form, type SubmissionHandler } from 'vee-validate';
import { useRouter } from 'vue-router';

import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';
import { useMutation } from '@/funcs/composable/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import { Register, type RegisterResponse } from '@/graphQL/Auth/auth';
import { useUserStore } from '@/stores/userStore';
import {
  FormKeys,
  type FormValues,
  initialValues,
  validationSchema,
} from '@/views/Auth/Register/form/RegisterForm';

const toast = useToast();
const router = useRouter();
const userStore = useUserStore();

const { execute } = useMutation<RegisterResponse>(Register, {
  isUseSpinner: true,
});

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (values: FormValues) => {
  await execute({
    input: {
      name: values[FormKeys.NAME],
      email: values[FormKeys.MAIL],
      password: values[FormKeys.PASSWORD],
    },
  }).then((res: RegisterResponse) => {
    toast.showToast({ message: '登録が完了しました' });
    //ログイン処理
    userStore.setUserData(res.registerUser);
    router.push({ name: 'Index' });
  });
};
</script>
<style scoped></style>
