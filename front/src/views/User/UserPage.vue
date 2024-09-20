<template>
  <UserData v-if="user" :user="user" />
</template>
<script setup lang="ts">
import { useRoute } from 'vue-router';
import { ref, watch } from 'vue';
import useQuery from '@/funcs/composable/useQuery';
import {
  GET_USERDATA_FULL,
  type GetUserByIdResponse,
  type User,
} from '@/graphQL/User/user';
import UserData from '@/components/templates/userPage/UserData.vue';

const user = ref<User>();

const route = useRoute();

const { fetch } = useQuery<GetUserByIdResponse>(GET_USERDATA_FULL);

// データフェッチ
const fetchData = async (id: string): Promise<void> => {
  const { getUserByIdDetail: response } = await fetch({
    id: id,
  });
  user.value = response;
};

watch(
  () => route.params.id, // ルートのパスやクエリ、パラメータなどを監視
  (to) => {
    // ルートが変更された際に実行される処理
    const id = to as string; // ルートパラメータからidを取得
    if (!id) {
      return;
    }
    fetchData(id);
  },
  { immediate: true }, // 初回レンダリング時に実行される
);
</script>

<style scoped></style>
