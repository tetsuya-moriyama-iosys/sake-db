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
    <UploadWithImage
      :name="FormKeys.IMAGE"
      :default="props.user.imageBase64"
      @onCompressed="onCompressed"
    />
    <FormField :name="FormKeys.PROFILE" label="プロフィール" as="textarea" />
    <SubmitButton>更新</SubmitButton>
  </Form>
</template>

<script setup lang="ts">
import { Form, type SubmissionHandler } from 'vee-validate';

import UploadWithImage from '@/components/parts/forms/common/UploadWithImage.vue';
import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';
import { useMutation } from '@/funcs/composable/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import { type AuthUser, type AuthUserFull, Update } from '@/graphQL/Auth/auth';
import { useUserStore } from '@/stores/userStore';
import {
  FormKeys,
  type FormValues,
  generateInitialValues,
  validationSchema,
} from '@/views/MyPage/UserEdit/form/EditForm';

interface Props {
  user: AuthUserFull;
}

const props = defineProps<Props>();

const toast = useToast();
const userStore = useUserStore();

const { execute } = useMutation<AuthUser>(Update, {
  isUseSpinner: true,
  isAuth: true,
});

let base64ImageData: string | undefined = props.user.imageBase64;

function onCompressed(encodedStr: string | null): void {
  base64ImageData = encodedStr ?? undefined;
}

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (values: FormValues) => {
  await execute({
    input: {
      name: values[FormKeys.NAME],
      email: values[FormKeys.MAIL],
      password: values[FormKeys.PASSWORD],
      profile: values[FormKeys.PROFILE],
      imageBase64: base64ImageData ?? null,
    },
  }).then(() => {
    toast.showToast({ message: '登録が完了しました' });
    //ユーザー情報のリフレッシュ
    userStore.restoreUserData({
      isReFetch: true,
    });
  });
};
</script>

<style scoped></style>
