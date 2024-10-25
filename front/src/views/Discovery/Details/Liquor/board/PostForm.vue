<template>
  感想を投稿
  <form @submit="onSubmit">
    <div class="flex board-container">
      <div class="flex-1 text-area">
        <FormField :name="FormKeys.TEXT" label="本文" classes="w-full" />
      </div>
      <RatingButton :name="FormKeys.RATE" label="評価" ref="ratingButton" />
      <div class="submit flex items-center justify-center ml-1.5">
        <SubmitButton class="h-10">送信</SubmitButton>
      </div>
    </div>
  </form>
</template>

<script setup lang="ts">
import { useForm } from 'vee-validate';
import { onMounted, ref } from 'vue';

import RatingButton from '@/components/parts/forms/common/RatingButton.vue';
import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';
import { useMutation, useQuery } from '@/funcs/composable/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import {
  GetMyPostByLiquorId,
  type MyBoardResponse,
  Post,
} from '@/graphQL/Liquor/board';
import {
  FormKeys,
  type FormValues,
  validationSchema,
} from '@/views/Discovery/Details/Liquor/board/LiquorBoard';

interface Props {
  liquorId: string;
}

const { liquorId } = defineProps<Props>();
const ratingButton = ref<InstanceType<typeof RatingButton> | null>(null);

const { fetch } = useQuery<MyBoardResponse>(GetMyPostByLiquorId, {
  isAuth: true,
}); //現在投稿されているものを初期値として取得する用
const { execute } = useMutation<boolean>(Post, { isAuth: true });
const toast = useToast();

const { handleSubmit, resetForm } = useForm<FormValues>({
  validationSchema, // yupのバリデーションスキーマを適用
});

const emit = defineEmits(['onSubmit']);

onMounted(async (): Promise<void> => {
  const response: MyBoardResponse = await fetch({
    id: liquorId,
  });
  if (response.getMyBoard == null) return;
  resetForm({
    values: {
      [FormKeys.TEXT]: response.getMyBoard.text,
      [FormKeys.RATE]: response.getMyBoard.rate ?? undefined,
    },
  });
  ratingButton.value?.resetRating(response.getMyBoard.rate ?? undefined); //評価コンポーネント内部状態のリセット
});

// フォームの送信処理
const onSubmit = handleSubmit(async (values: FormValues) => {
  await execute({
    input: {
      liquorID: liquorId,
      text: values[FormKeys.TEXT],
      rate: values[FormKeys.RATE],
    },
  }).then(() => {
    toast.showToast({ message: '投稿しました' });
    emit('onSubmit'); //リロードのコールバック
  });
});
</script>

<style scoped>
div.board-container {
  max-width: 800px;
}
</style>
