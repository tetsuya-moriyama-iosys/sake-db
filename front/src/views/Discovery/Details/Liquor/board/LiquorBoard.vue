<template>
  <PostForm :liquor-id="props.liquorId" @onSubmit="forceReload" />
  <PostList v-if="posts != null" :posts="posts" />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

import useQuery from '@/funcs/composable/useQuery/useQuery';
import {
  type BoardResponse,
  GET_BOARD,
  type Post,
} from '@/graphQL/Liquor/board';
import PostForm from '@/views/Discovery/Details/Liquor/board/PostForm.vue';
import PostList from '@/views/Discovery/Details/Liquor/board/PostList.vue';

interface Props {
  liquorId: string;
}

const props = defineProps<Props>();

const posts = ref<Post[] | null>(null);

const { fetch } = useQuery<BoardResponse>(GET_BOARD);

async function fetchData(isForceReload: boolean) {
  const response = await fetch(
    {
      liquorId: props.liquorId,
    },
    {
      fetchPolicy: isForceReload ? 'no-cache' : undefined,
    },
  );
  posts.value = response.board;
}
function forceReload() {
  void fetchData(true);
}

onMounted(async () => {
  void fetchData(false);
});
</script>

<style scoped></style>
