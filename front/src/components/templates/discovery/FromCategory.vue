<template>
  <div id="random-pickup-area">
    <CardContainer>
      <LiquorCard v-for="liquor in liquors" :liquor="liquor" :key="liquor.id" />
    </CardContainer>
  </div>
</template>

<script setup lang="ts">
import type { Liquor } from '@/graphQL/Index/random';
import CardContainer from '@/components/parts/common/CardContainer.vue';
import LiquorCard from '@/components/blocks/index/LiquorCard.vue';

interface Props {
  liquors: Liquor[];
}

const { liquors } = defineProps<Props>();
</script>

<style scoped>
div#random-pickup-area {
  margin: auto;
  padding: 2em;
}
</style>
