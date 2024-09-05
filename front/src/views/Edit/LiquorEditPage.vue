<template>
  投稿する
  <LiquorEdit v-if="!isLoading" :initial-data="liquor" />
</template>

<script setup lang="ts">
import useQuery from '@/funcs/composable/useQuery';
import { onMounted, ref } from 'vue';
import {
  LIQUOR_DETAIL_FOR_EDIT,
  type LiquorForEdit,
  type LiquorResponse,
} from '@/graphQL/Liquor/liquor';
import { useRoute } from 'vue-router';
import LiquorEdit from '@/components/templates/post/LiquorEdit.vue';

const isLoading = ref<boolean>(true);
const liquor = ref<LiquorForEdit | null>(null);

const route = useRoute(); // 現在のルートを取得
const { fetch } = useQuery<LiquorResponse<LiquorForEdit>>(
  LIQUOR_DETAIL_FOR_EDIT,
);

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const id = route.params.id as string; // ルートパラメータからidを取得
  if (!id) {
    isLoading.value = false;
    return;
  }
  const { liquor: response } = await fetch({
    variables: {
      id: id,
    },
  });
  liquor.value = response;
  isLoading.value = false;
});
</script>
