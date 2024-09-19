<template>
  <PostForm :liquor-id="props.liquorId" @onSubmit="forceReload" />
  <PostList v-if="posts != null" :posts="posts" />
</template>

<script setup lang="ts">
import PostForm from '@/components/blocks/discovery/liquorDetail/board/PostForm.vue';
import PostList from '@/components/blocks/discovery/liquorDetail/board/PostList.vue';
import {
  type BoardResponse,
  GET_BOARD,
  type Post,
} from '@/graphQL/Liquor/board';
import { onMounted, ref } from 'vue';
import useQuery from '@/funcs/composable/useQuery';

interface Props {
  liquorId: string;
}

const props = defineProps<Props>();

const posts = ref<Post[] | null>(null);

const { fetch } = useQuery<BoardResponse>(GET_BOARD);

async function fetchData(isForceReload: boolean) {
  const response = await fetch({
    variables: {
      liquorId: props.liquorId,
    },
    fetchPolicy: isForceReload ? 'no-cache' : undefined,
  });
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
