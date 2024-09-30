<template>
  このユーザーをブックマークしているユーザー
  <span v-if="userList">({{ userList.length }}人)</span>
  <UserList
    v-if="userList"
    :user-list="userList"
    :is-show-created-at="false"
    :is-show-header="false"
  />
</template>

<script setup lang="ts">
import useQuery from '@/funcs/composable/useQuery';
import {
  type Bookmark,
  BOOKMARKED_LIST,
  type GetBookmarkedListResponse,
} from '@/graphQL/Bookmark/bookmark';
import { onMounted, ref } from 'vue';
import UserList from '@/components/blocks/common/userPage/UserList.vue';

const { fetch } = useQuery<GetBookmarkedListResponse>(BOOKMARKED_LIST);

const userList = ref<Bookmark[] | null>(null);

const props = defineProps<{
  id: string;
}>();

onMounted(async () => {
  const { getBookMarkedList: response } = await fetch({ id: props.id });
  userList.value = response ?? [];
});
</script>

<style scoped></style>
