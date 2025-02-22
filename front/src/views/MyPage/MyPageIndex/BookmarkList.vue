<template>
  ブックマークリスト
  <div v-if="bookmarks">
    <UserList
      v-if="bookmarks.length > 0"
      :is-show-created-at="true"
      :user-list="bookmarks"
      v-slot="{ user }"
    >
      <BookMarkLogics :target-id="user.userId" v-slot="{ remove }">
        <CommonButton size="small" @click="deleteUser(user.userId, remove)"
          >削除</CommonButton
        >
      </BookMarkLogics>
    </UserList>
  </div>
  <div v-else>ブックマークしているユーザーはいません</div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

import UserList from '@/components/blocks/common/userPage/UserList.vue';
import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import BookMarkLogics from '@/components/slots/BookMarkLogics.vue';
import useQuery from '@/funcs/composable/useQuery/useQuery';
import {
  type Bookmark,
  type GetBookmarkListResponse,
  LIST,
} from '@/graphQL/Bookmark/bookmark';

const { fetch } = useQuery<GetBookmarkListResponse>(LIST, {
  isAuth: true,
});
const bookmarks = ref<Bookmark[] | null>(null);

onMounted(() => {
  void reFetch({
    isUseCache: false, //誰かをブックマークしてからリストに戻っても反映されないのでキャッシュを使わないことにした
  });
});

const reFetch = async ({
  isUseCache = false,
}: {
  isUseCache: boolean;
}): Promise<void> => {
  const response = await fetch(undefined, {
    fetchPolicy: isUseCache ? 'cache-first' : 'no-cache',
  });
  bookmarks.value = response.getBookMarkList ?? [];
};

//削除ボタンの動作
const deleteUser = async (userId: string, removeFn: () => Promise<void>) => {
  await removeFn();
  //再度取得する
  void reFetch({
    isUseCache: false,
  });
};
</script>

<style scoped></style>
