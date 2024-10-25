<!--汎用ボタン-->
<template>
  <button :class="buttonClass" :disabled="isDisabled">
    <slot />
    <div
      v-if="!isDisabled"
      class="absolute inset-0 flex h-full w-full justify-center [transform:skew(-12deg)_translateX(-100%)] group-hover:duration-1000 group-hover:[transform:skew(-12deg)_translateX(100%)]"
    >
      <div class="relative h-full w-8 bg-white/20"></div>
    </div>
  </button>
</template>

<script setup lang="ts">
import { computed } from 'vue';

import type { ButtonProps } from '@/components/parts/common/CommonButton/type';
import { ColorType } from '@/type/common/ColorType';

const props = defineProps<ButtonProps>();

const BaseButtonClass: string =
  'group relative inline-flex items-center justify-center overflow-hidden rounded-md px-6 font-medium';
const EnabledClass: string = `${BaseButtonClass} transition hover:scale-110 text-white`;
const DisabledClass: string = `${BaseButtonClass} text-neutral-400 cursor-not-allowed`;

// ボタンのクラスを動的に決定するコンピューテッドプロパティ
const buttonClass = computed(() => {
  const baseClass = `${props.isDisabled ? DisabledClass : EnabledClass} ${props.class} ${buttonSize.value}`;

  if (props.isDisabled) {
    switch (props.color) {
      case ColorType.Primary:
        return `${baseClass} bg-blue-300`;
      case ColorType.Secondary:
        return `${baseClass} bg-green-300`;
      case ColorType.Danger:
        return `${baseClass} bg-red-300`;
      default:
        return `${baseClass} bg-neutral-300`;
    }
  } else {
    switch (props.color) {
      case ColorType.Primary:
        return `${baseClass} bg-blue-500 hover:bg-blue-600`;
      case ColorType.Secondary:
        return `${baseClass} bg-green-500 hover:bg-green-600`;
      case ColorType.Danger:
        return `${baseClass} bg-red-500 hover:bg-red-600`;
      default:
        return `${baseClass} bg-neutral-950 hover:bg-neutral-800`;
    }
  }
});

const buttonSize = computed(() => {
  switch (props.size) {
    case 'small':
      return 'h-8';
    case 'large':
      return 'h-16';
    default:
      return 'h-12'; // デフォルト
  }
});
</script>

<style scoped></style>
