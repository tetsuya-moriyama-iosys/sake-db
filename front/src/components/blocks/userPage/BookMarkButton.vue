<template>
  <BookMarkLogics
    :target-id="props.targetId"
    v-slot="{ isBookmarked, onClick }"
    :default-state="undefined"
  >
    <CommonButton
      v-if="!isBookmarked"
      @click="onClick"
      :color="ColorType.Secondary"
    >
      ブックマーク
    </CommonButton>
    <CommonButton v-else @click="onClick" :color="ColorType.Danger">
      ブックマーク済
    </CommonButton>
  </BookMarkLogics>
</template>

<script setup lang="ts">
import BookMarkLogics from '@/components/slots/BookMarkLogics.vue';
import CommonButton from '@/components/parts/common/CommonButton.vue';
import { ColorType } from '@/type/common/ColorType';

const props = defineProps<{
  targetId: string;
}>();
</script>

<style scoped></style>
