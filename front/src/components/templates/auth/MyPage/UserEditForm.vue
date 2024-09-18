<template>
  <Form
    @submit="onSubmit"
    :initial-values="generateInitialValues(props.user)"
    :validation-schema="validationSchema"
  >
    <FormField :name="FormKeys.NAME" label="名前" />
    <FormField :name="FormKeys.MAIL" label="メールアドレス" type="email" />
    <FormField
      :name="FormKeys.PASSWORD"
      label="パスワード(変更する場合のみ)"
      type="password"
    />
    <UploadWithImage :name="FormKeys.IMAGE" :default="props.user.imageBase64" />
    <FormField :name="FormKeys.PROFILE" label="プロフィール" as="textarea" />
  </Form>
</template>
<script setup lang="ts">
import { Form, type SubmissionHandler } from 'vee-validate';
import FormField from '@/components/parts/forms/core/FormField.vue';
import { useToast } from '@/funcs/composable/useToast';
import { useMutation } from '@/funcs/composable/useQuery';
import {
  FormKeys,
  type FormValues,
  generateInitialValues,
  validationSchema,
} from '@/forms/define/auth/EditForm';
import type { AuthUser, UserFullData } from '@/graphQL/User/user';
import { Register } from '@/graphQL/Auth/auth';
import UploadWithImage from '@/components/parts/forms/common/UploadWithImage.vue';

interface Props {
  user: UserFullData;
}

const props = defineProps<Props>();

const toast = useToast();
const { execute } = useMutation<AuthUser>(Register, {
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
  }).then(() => {
    toast.showToast({ message: '登録が完了しました' });
    //ヘッダー等の更新処理
  });
};
</script>

<style scoped></style>
