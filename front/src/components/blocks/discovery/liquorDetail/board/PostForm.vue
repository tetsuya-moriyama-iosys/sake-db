<template>
  感想を投稿
  <VForm @submit="onSubmit" :validation-schema="validationSchema" ref="form">
    <div class="flex board-container">
      <div class="flex-1 text-area">
        <FormField :name="FormKeys.TEXT" label="本文" classes="w-full" />
      </div>
      <RatingButton :name="FormKeys.RATE" label="評価" ref="ratingButton" />
      <div class="submit flex items-center justify-center ml-1.5">
        <SubmitButton class="h-10">送信</SubmitButton>
      </div>
    </div>
  </VForm>
</template>

<script setup lang="ts">
import { Form as VForm, type SubmissionHandler } from 'vee-validate';
import FormField from '@/components/parts/forms/core/FormField.vue';
import RatingButton from '@/components/parts/forms/common/RatingButton.vue';
import {
  FormKeys,
  type FormValues,
  validationSchema,
} from '@/forms/define/details/LiquorBoard';
import SubmitButton from '@/components/parts/common/SubmitButton.vue';
import { useMutation } from '@/funcs/composable/useQuery';
import { Post } from '@/graphQL/Liquor/board';
import { ref } from 'vue';

interface Props {
  liquorId: string;
}

const { liquorId } = defineProps<Props>();
const form = ref<InstanceType<typeof VForm> | null>(null); //Form内部に定義されているフォームメソッドにアクセスするのに必要
const ratingButton = ref<InstanceType<typeof RatingButton> | null>(null);

const { execute } = useMutation<boolean>(Post, { isAuth: true });

const emit = defineEmits(['onSubmit']);

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
async function onSubmit(values: FormValues): SubmissionHandler {
  console.log('values:', values);
  await execute({
    variables: {
      input: {
        liquorID: liquorId,
        text: values[FormKeys.TEXT],
        rate: values[FormKeys.RATE],
      },
    },
  }).then(() => {
    form.value?.resetForm();
    ratingButton.value?.resetRating(); //評価コンポーネント内部状態のリセット
    emit('onSubmit'); //リロードのコールバック
  });
}
</script>

<style scoped>
div.board-container {
  max-width: 800px;
}
</style>
