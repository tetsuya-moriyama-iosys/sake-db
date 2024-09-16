<!--酒の情報ページ-->
<template>
  <LiquorDetail v-if="liquor" :liquor="liquor" />
</template>

<script setup lang="ts">
import useQuery from '@/funcs/composable/useQuery';
import { ref, watch } from 'vue';
import LiquorDetail from '@/components/templates/discovery/details/liquor/LiquorDetail.vue';
import {
  type Liquor,
  LIQUOR_DETAIL_GET,
  type LiquorResponse,
} from '@/graphQL/Liquor/liquor';
import { useRoute } from 'vue-router';
import { useSelectedCategoryStore } from '@/stores/sidebar';

const isLoading = ref<boolean>(true);

const route = useRoute(); // 現在のルートを取得
const sidebarStore = useSelectedCategoryStore();
const { fetch } = useQuery<LiquorResponse<Liquor>>(LIQUOR_DETAIL_GET);

const liquor = ref<Liquor | null>(null);

const isNoCache: boolean = window.history.state?.noCache ?? false; //TODO:何故か常にtrueになってる...？

// データフェッチ
const fetchData = async (id: string): Promise<void> => {
  const { liquor: response } = await fetch({
    variables: {
      id: id,
    },
    fetchPolicy: isNoCache ? 'no-cache' : undefined, //更新直後だとキャッシュが残っているため、キャッシュを無効化
  });
  liquor.value = response;
  sidebarStore.updateContent(response.categoryId);
  isLoading.value = false;
};

watch(
  () => route.params.id, // ルートのパスやクエリ、パラメータなどを監視
  (to) => {
    // ルートが変更された際に実行される処理
    const id = to as string; // ルートパラメータからidを取得
    if (!id) {
      isLoading.value = false;
      return;
    }
    fetchData(id);
  },
  { immediate: true }, // 初回レンダリング時に実行される
);
</script>
