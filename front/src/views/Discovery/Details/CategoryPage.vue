<template>
  <CategoryDetail v-if="category" :category="category" />
</template>
<script setup lang="ts">
import { useRoute } from 'vue-router';
import useQuery from '@/funcs/composable/useQuery';
import {
  type Category,
  type CategoryResponse,
  GET_DETAIL,
} from '@/graphQL/Category/categories';
import { onMounted, ref } from 'vue';
import CategoryDetail from '@/components/templates/discovery/details/CategoryDetail.vue';

const route = useRoute(); // 現在のルートを取得
const { fetch } = useQuery<CategoryResponse<Category>>(GET_DETAIL);

const category = ref<Category | null>(null);

const isNoCache: boolean = window.history.state?.noCache ?? false; //TODO:何故か常にtrueになってる...？

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const { category: response } = await fetch({
    variables: {
      id: route.params.id, // ルートパラメータからidを取得,
    },
    fetchPolicy: isNoCache ? 'no-cache' : undefined, //更新直後だとキャッシュが残っているため、キャッシュを無効化
  });
  category.value = response;
});
</script>

<style scoped></style>
