<template>
  <span class="tag"
    >{{ props.text }}
    <div v-if="props.isClose" class="close inline-block" @click="emit('close')">
      <span
        class="text-gray-500 transition-colors duration-300 ease-in-out"
      ></span></div
  ></span>
</template>

<script setup lang="ts">
import type { TagProps } from '@/components/parts/common/CommonTag/type';

const props = defineProps<TagProps>();

const emit = defineEmits(['close']);
</script>

<style scoped>
span.tag {
  display: inline-block;
  margin: 0 0.1em 0.6em 0;
  padding: 0.6em;
  line-height: 1;
  text-decoration: none;
  color: #0000ee;
  background-color: #fff;
  border: 1px solid #0000ee;
  border-radius: 2em;
  cursor: pointer;

  &:before {
    content: '#';
  }
}

div.close {
  span {
    display: inline-block;
    vertical-align: middle;
    color: #333;
    line-height: 1;
    width: 1em;
    height: 0.1em;
    background: currentColor;
    border-radius: 0.1em;
    position: relative;
    transform: rotate(45deg);

    &::before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background: inherit;
      border-radius: inherit;
      transform: rotate(90deg);
    }
  }

  &:hover span {
    color: #ff0000; /* ホバー時に色を赤に変更 */
  }
}
</style>
