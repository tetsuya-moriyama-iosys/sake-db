<template>
  <div
    class="image-container"
    :style="{
      width: computedSize,
      height: computedSize,
      borderRadius: radius,
    }"
  >
    <img v-if="imageSrc" :src="imageSrc" alt="画像" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue';
import convertToBase64 from '@/funcs/util/convertToBase64 ';

interface Props {
  imageSrc?: string | File; // Base64エンコードされた文字列
  radius?: string; // 任意の半径 (px, %, emなど)
}

const props = defineProps<Props>();

// 画像のsrcを管理するRefオブジェクト
const imageSrc = ref<string | undefined>(
  typeof props.imageSrc === 'string' ? props.imageSrc : undefined,
);

// デフォルトの値
const radius = props.radius ?? '50px'; // デフォルトの半径は50px

// 半径を元に幅と高さを動的に計算 (半径 × 2)
const computedSize = computed(() => {
  const numRadius = parseInt(radius, 10); // 数値に変換
  if (isNaN(numRadius)) return '100px'; // radiusが不正な場合のデフォルト値
  return `${numRadius * 2}px`;
});

// Fileオブジェクトが変更されたらBase64に変換してimageSrcにセットする
watch(
  () => props.imageSrc,
  async (newFile: string | File | undefined) => {
    if (!newFile) return;
    try {
      //refを更新する
      imageSrc.value =
        typeof newFile === 'string' ? newFile : await convertToBase64(newFile);
    } catch (error) {
      console.error('ファイルの変換に失敗しました', error);
    }
  },
);
</script>

<style scoped>
.image-container {
  overflow: hidden; /* 丸型に切り抜くために必要 */
}

.image-container img {
  width: 100%;
  height: 100%;
  object-fit: cover; /* 画像の比率を保ちながらコンテナにフィットさせる */
}
</style>
