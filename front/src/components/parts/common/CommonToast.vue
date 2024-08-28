<template>
  <Transition name="toast">
    <div
      v-show="isVisible"
      :class="[
        'fixed top-4 left-1/2 transform -translate-x-1/2 z-50 p-4 rounded-md shadow-lg',
        toastColorClasses,
      ]"
    >
      {{ savedProps?.message }}
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue';
import { type ToastOptions, toastState, ToastType } from '@/plugins/toast';

const isVisible = ref<boolean>(false); //DOM要素についてのstate
const savedProps = ref<ToastOptions | null>(null); //元のオブジェクトを参照すると、消去時にnullに書き換わってしまうので状態を保存する必要がある

const toastColorClasses = computed(() => {
  switch (savedProps.value?.type) {
    case ToastType.Success:
      return 'bg-green-500';
    case ToastType.Error:
      return 'bg-red-500';
    case ToastType.Info:
    default:
      return 'bg-blue-500';
  }
});

// トーストの状態を監視
watch(
  () => toastState.value,
  (newState: ToastOptions | null) => {
    if (newState === null) {
      return;
    }
    savedProps.value = newState;
    isVisible.value = true;
    setTimeout(() => {
      isVisible.value = false;
    }, newState.duration);
  },
);
</script>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition: opacity 0.3s ease;
}

.toast-enter-from,
.toast-leave-to {
  opacity: 0;
}
</style>
