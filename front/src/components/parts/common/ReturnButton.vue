<template>
  <CommonButton
    v-bind="buttonProps"
    @click="
      () => {
        if (to == null) {
          router.back();
          return;
        }
        router.push({ name: to, params });
      }
    "
    >戻る</CommonButton
  >
</template>

<script setup lang="ts">
import type { ButtonProps } from '@/components/parts/common/CommonButton/type';

const router = useRouter();

const { to, params, ...buttonProps } = defineProps<
  ButtonProps & {
    to?: string;
    params?: Record<string, string>;
  }
>();

import { useRouter } from 'vue-router';

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
</script>
