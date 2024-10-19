<template>
  <MenuItem v-slot="{ close }">
    <router-link :to="{ name: props.linkTo }"
      ><div
        @click="handleClick(close)"
        class="block px-4 py-2 text-gray-800 hover:bg-gray-100"
      >
        <slot /></div
    ></router-link>
  </MenuItem>
</template>

<script setup lang="ts">
import { MenuItem } from '@headlessui/vue';

// イベントのemit関数を定義
const emit = defineEmits(['click']);

const props = defineProps<{
  linkTo: string;
}>();

const handleClick = (close: () => void) => {
  // 親コンポーネントから渡された@clickイベントをemitで発火
  emit('click');
  close(); // メニューを閉じる
};
</script>

<style scoped></style>
