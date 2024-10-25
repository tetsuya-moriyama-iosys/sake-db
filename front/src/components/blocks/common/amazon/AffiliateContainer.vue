<template>
  <div v-if="affiliateData" class="affiliate-container m-auto">
    <CardContainer>
      <AffiliateCard
        v-for="item in affiliateData.items"
        :item="item"
        :key="item.URL"
      />
    </CardContainer>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

import AffiliateCard from '@/components/blocks/common/amazon/AffiliateCard.vue';
import CardContainer from '@/components/parts/common/CardContainer.vue';
import useQuery from '@/funcs/composable/useQuery';
import {
  type AffiliateData,
  type AffiliateResponse,
  GET_AFFILIATE_LIST,
} from '@/graphQL/Amazon/affiliate';

interface Props {
  name: string; //商品名
  limit?: number;
}
const { name, limit } = defineProps<Props>();

const { fetch } = useQuery<AffiliateResponse>(GET_AFFILIATE_LIST);

const affiliateData = ref<AffiliateData | null>(null);

onMounted(async () => {
  const { data: response } = await fetch({
    keyword: name,
    limit: limit ?? 5,
  });
  affiliateData.value = response;
});
</script>

<style scoped>
div.affiliate-container {
  max-width: 1250px;
}
</style>
