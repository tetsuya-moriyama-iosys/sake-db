<template>
  <UserData v-if="userDetail" :userDetail="userDetail" />
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';

import useQuery from '@/funcs/composable/useQuery/useQuery';
import {
  GET_USERDATA_FULL,
  type GetUserDetailResponse,
  type UserDetail,
} from '@/graphQL/User/user';
import UserData from '@/views/User/UserData.vue';

const userDetail = ref<UserDetail>();

const route = useRoute();

const { fetch } = useQuery<GetUserDetailResponse>(GET_USERDATA_FULL);

// データフェッチ
const fetchData = async (id: string): Promise<void> => {
  const { getUserByIdDetail: response } = await fetch(
    {
      id: id,
    },
    {
      fetchPolicy: 'no-cache', // なぜかこれを付けないとrateがnullになる(キャッシュが原因なのか...？)
    },
  );
  userDetail.value = response;
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
