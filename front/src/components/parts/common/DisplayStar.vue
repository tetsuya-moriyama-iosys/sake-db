<template>
  <span v-for="index in props.rate" :key="index">★</span>
</template>

<script setup lang="ts">
const props = defineProps<{
  rate: number | null;
}>();
</script>
