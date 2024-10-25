<template>
  <CommonDialog v-bind="props" is-un-use-default-buttons v-slot="{ close }">
    <div><slot /></div>
    <div class="mt-5 flex justify-center gap-4">
      <CommonButton
        size="small"
        @click="
          async () => {
            await handleYes().finally(() => {
              close();
            });
          }
        "
        >{{ props.yes }}</CommonButton
      >
      <CommonButton
        size="small"
        @click="
          async () => {
            await handleNo().finally(() => {
              close();
            });
          }
        "
        >{{ props.no }}</CommonButton
      >
    </div>
  </CommonDialog>
</template>

<script setup lang="ts">
import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import CommonDialog from '@/components/parts/common/CommonDialog/CommonDialog.vue';
import type { YesNoDialogProps } from '@/components/parts/common/CommonDialog/Variations/type';

const props = withDefaults(defineProps<YesNoDialogProps>(), {
  yes: 'はい',
  no: 'いいえ',
});

async function handleYes() {
  return props.onYes();
}
async function handleNo() {
  return props.onNo?.();
}
</script>

<style scoped></style>
