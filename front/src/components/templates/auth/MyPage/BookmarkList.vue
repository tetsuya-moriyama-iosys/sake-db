<template>
  ブックマークリスト
  <table>
    <tr>
      <td>ユーザー名</td>
      <td>追加日</td>
      <td></td>
    </tr>
    <tr v-for="bookmark in bookmarks" :key="bookmark.userId">
      <td>
        <router-link
          :to="{ name: 'UserPage', params: { id: bookmark.userId } }"
          >{{ bookmark.name }}</router-link
        >
      </td>
      <td>{{ format(date(bookmark.createdAt), 'yyyy/MM/dd') }}</td>
      <td>
        <CommonButton size="small">削除</CommonButton>
      </td>
    </tr>
  </table>
</template>

<script setup lang="ts">
import {
  type Bookmark,
  type GetBookmarkListResponse,
  LIST,
} from '@/graphQL/Bookmark/bookmark';
import useQuery from '@/funcs/composable/useQuery';
import { onMounted, ref } from 'vue';
import date from '@/funcs/util/date';
import { format } from 'date-fns';
import CommonButton from '@/components/parts/common/CommonButton.vue';

const { fetch } = useQuery<GetBookmarkListResponse>(LIST, {
  isAuth: true,
});
const bookmarks = ref<Bookmark[] | null>(null);

onMounted(async () => {
  const response = await fetch();
  bookmarks.value = response.getBookMarkList ?? [];
});
</script>

<style scoped></style>
