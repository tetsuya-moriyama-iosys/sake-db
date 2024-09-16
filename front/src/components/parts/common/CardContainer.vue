<template>
  <div class="card-container" :style="containerStyles">
    <slot />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

interface Props {
  columns?: string;
  gap?: string;
  min?: string;
  isUseAutoFit?: boolean;
}

// Props の受け取り
const props = withDefaults(defineProps<Props>(), {
  columns: 'repeat(auto-fill, minmax(200px, 1fr)',
  gap: '1em',
  min: '200px',
});

// 動的なスタイルの定義
const containerStyles = computed(() => ({
  gridTemplateColumns: `repeat(${props.isUseAutoFit ? 'auto-fit' : 'auto-fill'}, minmax(${props.min}, 1fr))`,
  gap: props.gap,
}));
</script>

<style scoped>
div.card-container {
  display: grid;
  width: 100%;
}
</style>
