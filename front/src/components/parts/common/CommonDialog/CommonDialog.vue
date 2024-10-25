<!--モーダルコンポーネント。親のrefで開閉状態を管理する。slotで実装するか迷ったが、状態の取り回しがまどろっこしいのでrefは親コンポーネントで定義してもらう。-->
<template>
  <Dialog :open="isOpen" @close="handleClose">
    <div
      class="fixed flex inset-0 bg-black/30 overflow-y-auto items-center justify-center"
      :aria-hidden="!isOpen"
    >
      <div
        class="dialog-container bg-white p-4 text-center rounded-2xl transform -translate-y-52"
        :class="props.class"
      >
        <DialogPanel>
          <DialogTitle
            v-if="props.title"
            as="h3"
            class="text-lg font-medium leading-6 text-gray-900"
            >{{ props.title }}</DialogTitle
          >
          <slot :close="handleClose"></slot>
          <div v-if="!props.isUnUseDefaultButtons">
            <CommonButton @click="handleClose">閉じる</CommonButton>
          </div>
        </DialogPanel>
      </div>
    </div>
  </Dialog>
</template>

<script setup lang="ts">
import { Dialog, DialogPanel, DialogTitle } from '@headlessui/vue';

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import type { DialogProps } from '@/components/parts/common/CommonDialog/type';

const props = defineProps<DialogProps>();
const isOpen = defineModel<boolean>();

function handleClose() {
  isOpen.value = false;
}
</script>

<style scoped>
div.dialog-container {
  min-width: 30em;
}
</style>
