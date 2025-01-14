<template>
  <CommonButton @click="showModal">タグ入力</CommonButton>
  <CommonDialog
    title="タグ入力"
    v-model="isDialogOpen"
    is-un-use-default-buttons
    v-slot="{ close }"
  >
    <!--submitの警告は、vee-validateのGenericObjectの型が広すぎることによるので無視してOK-->
    <Form
      :initial-values="defaultValues(props.liquorId)"
      :validation-schema="validationSchema"
      @submit="
        async (values: FormValues) => {
          await onSubmit(values);
          close();
        }
      "
    >
      <FormField :name="PostTagKeys.LiquorId" type="hidden" />
      <FormField :name="PostTagKeys.Tag" />
      <div>
        <SubmitButton size="small">登録</SubmitButton>
        <CommonButton size="small" @click="close">閉じる</CommonButton>
      </div>
    </Form>
  </CommonDialog>
</template>

<script setup lang="ts">
import { Form } from 'vee-validate';
import { ref } from 'vue';

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import CommonDialog from '@/components/parts/common/CommonDialog/CommonDialog.vue';
import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';
import { useMutation } from '@/funcs/composable/useQuery/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import {
  PostTag,
  PostTagKeys,
  type PostTagResponse,
} from '@/graphQL/Liquor/tags';

import { defaultValues, type FormValues, validationSchema } from './form';

const props = defineProps<{
  liquorId: string;
}>();

const { execute } = useMutation<PostTagResponse>(PostTag, {
  isAuth: true,
});
const toast = useToast();

const emit = defineEmits(['submitted']);

const isDialogOpen = ref<boolean>(false);

async function onSubmit(values: FormValues) {
  const response = await execute({ input: values });
  toast.showToast({
    message: 'タグの登録に成功しました',
  });
  emit('submitted', response.postTag);
}

function showModal() {
  isDialogOpen.value = true;
}
</script>

<style scoped></style>
