<!--酒の情報ページ-->
<template>
  <LiquorDetail v-if="liquor" :liquor="liquor" />
</template>

<script setup lang="ts">
import useQuery from '@/funcs/composable/useQuery';
import { onMounted, ref } from 'vue';
import LiquorDetail from '@/components/templates/discovery/LiquorDetail.vue';
import {
  type Liquor,
  LIQUOR_DETAIL_GET,
  type LiquorResponse,
} from '@/graphQL/Liquor/liquor';
import { useRoute } from 'vue-router';

const isLoading = ref<boolean>(true);

const route = useRoute(); // 現在のルートを取得
const { fetch } = useQuery<LiquorResponse<Liquor>>(LIQUOR_DETAIL_GET);

const liquor = ref<Liquor | null>(null);

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
