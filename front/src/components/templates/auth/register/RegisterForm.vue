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
import FormField from '@/components/parts/forms/core/FormField.vue';
import {
  FormKeys,
  type FormValues,
  initialValues,
  validationSchema,
} from '@/forms/auth/RegisterForm';
import SubmitButton from '@/components/parts/common/SubmitButton.vue';
import { REGISTER, type User } from '@/graphQL/Auth/register';
import { useMutation } from '@/funcs/composable/useQuery';

const { execute } = useMutation<User>(REGISTER, {
  isUseSpinner: true,
});

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (values: FormValues) => {
  await execute({
    variables: {
      input: {
        name: values[FormKeys.NAME],
        email: values[FormKeys.MAIL],
        password: values[FormKeys.PASSWORD],
      },
    },
  });
};
</script>
<style scoped></style>
