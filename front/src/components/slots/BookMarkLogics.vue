<template>
  <slot :isBookmarked="isBookmarked" :onClick="onClick" :remove="remove" />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

import { useMutation, useQuery } from '@/funcs/composable/useQuery/useQuery';
import {
  ADD,
  type AddResponse,
  CHECK,
  type CheckResponse,
  REMOVE,
  type RemoveResponse,
} from '@/graphQL/Bookmark/bookmark';
import { useUserStore } from '@/stores/userStore/userStore';

const userStore = useUserStore();
const isBookmarked = ref<boolean>(false);

const props = defineProps<{
  targetId: string;
  defaultState?: boolean; //APIを介して初期設定する必要がない場合は設定する(memo:なぜか明示的にundefinedを渡さないとfalseになってるバグがある？)
}>();

const { fetch } = useQuery<CheckResponse>(CHECK, {
  isAuth: true,
});
const { execute: addExecute } = useMutation<AddResponse>(ADD, {
  isAuth: true,
});
const { execute: removeExecute } = useMutation<RemoveResponse>(REMOVE, {
  isAuth: true,
});

onMounted(async (): Promise<void> => {
  //なぜか型推論がおかしいのでキャスト(booleanとして扱われてしまう)
  if ((props.defaultState as boolean | undefined) !== undefined) {
    //初期値が与えられていた場合、それにセットして終了
    isBookmarked.value = props.defaultState;
    return;
  }
  //初期値がない場合、APIから取得する
  const response = await fetch(
    {
      id: props.targetId,
    },
    {
      fetchPolicy: 'no-cache',
    },
  );
  isBookmarked.value = response.getIsBookMarked;
});

async function onClick(): Promise<void> {
  //未ログインなら終了
  if (!userStore.isLogin) {
    return;
  }
  //自分自身を追加しない
  if (userStore.user?.id === props.targetId) {
    return;
  }
  if (isBookmarked.value) {
    await remove();
  } else {
    await add();
  }
}
async function add(): Promise<void> {
  await addExecute({
    id: props.targetId,
  });
  isBookmarked.value = true;
}

async function remove(): Promise<void> {
  await removeExecute({
    id: props.targetId,
  });
  isBookmarked.value = false;
}
</script>
