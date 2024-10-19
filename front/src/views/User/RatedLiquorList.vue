<template>
  <div v-if="props.liquorList != null && props.liquorList.length != 0">
    <CardContainer min="150px">
      <UserLiquorCard
        v-for="liquor in liquorList"
        :liquor="liquor"
        :key="liquor.liquorId"
      />
    </CardContainer>
  </div>
  <div v-else>まだ評価がありません</div>
</template>
<script setup lang="ts">
import CardContainer from '@/components/parts/common/CardContainer.vue';
import type { UserLiquor } from '@/graphQL/User/user';
import UserLiquorCard from '@/views/User/UserLiquorCard.vue';

interface Props {
  liquorList: UserLiquor[] | null; //内容が空の場合の表示制御もこのコンポーネント内でやる想定
}
const props = defineProps<Props>();
</script>
<style scoped></style>
