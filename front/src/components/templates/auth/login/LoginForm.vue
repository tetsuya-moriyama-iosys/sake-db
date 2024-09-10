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
</template>

<script setup lang="ts">
import { Form, type SubmissionHandler } from 'vee-validate';
import FormField from '@/components/parts/forms/core/FormField.vue';
import {
  FormKeys,
  type FormValues,
  initialValues,
  validationSchema,
} from '@/forms/auth/LoginForm';
import SubmitButton from '@/components/parts/common/SubmitButton.vue';
import { LOGIN, type User } from '@/graphQL/Auth/register';
import { useMutation } from '@/funcs/composable/useQuery';

const { execute } = useMutation<User>(LOGIN);

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (values: FormValues) => {
  await execute({
    variables: {
      input: {
        email: values[FormKeys.MAIL],
        password: values[FormKeys.PASSWORD],
      },
    },
  });
};
</script>
<style scoped></style>
