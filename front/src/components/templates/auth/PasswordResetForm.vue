<template>
  <Form
    @submit="onSubmit"
    :validation-schema="{
      [EMAIL_NAME]: string().required().email(),
    }"
  >
    <FormField name="email" label="メールアドレス" type="email" />
    <SubmitButton />
  </Form>
</template>

<script setup lang="ts">
import { Form, type SubmissionHandler } from 'vee-validate';
import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/common/SubmitButton.vue';
import { string } from 'yup';
import { useToast } from '@/funcs/composable/useToast';
import { useMutation } from '@/funcs/composable/useQuery';
import { PASSWORD_RESET } from '@/graphQL/Auth/auth';
import { ToastType } from '@/plugins/toast';

//メールアドレス以外現状要らないので、定数もここで定義
const EMAIL_NAME = 'email';

const toast = useToast();

const { execute } = useMutation<{
  passwordReset: boolean;
}>(PASSWORD_RESET, {
  isUseSpinner: true,
});

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (values: {
  [EMAIL_NAME]: string;
}) => {
  await execute({
    email: values[EMAIL_NAME],
  }).then((res) => {
    if (res.passwordReset) {
      toast.showToast({
        message: 'パスワードリセットメールを送信しました。',
        type: ToastType.Success,
      });
    }
  });
};
</script>

<style scoped></style>
