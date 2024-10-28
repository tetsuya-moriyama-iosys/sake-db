<template>
  <div
    class="grid-cell"
    :style="{ backgroundColor: getBackgroundColor(props.cellData.rate) }"
  ></div>
</template>

<script setup lang="ts">
import type { FlavorCell } from '@/graphQL/Liquor/flavorMap';

const props = defineProps<{
  isSelected: boolean;
  cellData: FlavorCell;
}>();

// rate に基づいて背景色を決定する関数
function getBackgroundColor(rate: number): string {
  if (props.isSelected) {
    return `rgba(0, 0, 255, 1)`;
  }
  // 50%以上の得票率で最大になるように調整
  const opacity = rate >= 50 ? 1 : Math.min(Math.max(rate / 50, 0), 1) * 0.5;
  return `rgba(255, 0, 0, ${opacity})`; // 赤色の濃淡を調整（適宜変更）
}
</script>

<style scoped lang="scss">
div.grid-cell {
  width: calc(420px / 21);
  transition: background-color 0.3s; // 背景色の変化にトランジションを適用 (任意)

  &:hover {
    background-color: #666666 !important; // styleより優先度を上げるためimportant
  }
}
</style>
