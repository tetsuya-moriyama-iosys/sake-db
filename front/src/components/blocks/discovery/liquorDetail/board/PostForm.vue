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
import { useMutation, useQuery } from '@/funcs/composable/useQuery';
import {
  type BoardResponse,
  GetMyPostByLiquorId,
  myBoardRequest,
  Post,
  type PostCore,
} from '@/graphQL/Liquor/board';
import { onMounted, ref } from 'vue';

interface Props {
  liquorId: string;
}

const { liquorId } = defineProps<Props>();
const form = ref<InstanceType<typeof VForm> | null>(null); //Form内部に定義されているフォームメソッドにアクセスするのに必要
const ratingButton = ref<InstanceType<typeof RatingButton> | null>(null);

const { fetch } = useQuery<BoardResponse<PostCore>>(GetMyPostByLiquorId, {
  isAuth: true,
}); //現在投稿されているものを初期値として取得する用
const { execute } = useMutation<boolean>(Post, { isAuth: true });

const emit = defineEmits(['onSubmit']);

onMounted(async (): Promise<void> => {
  const response: BoardResponse<PostCore> | undefined = await fetch(
    myBoardRequest(liquorId),
  );
  console.log('response:', response);
});

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
