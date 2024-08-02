<template>
  <HeaderArea />
  <router-view></router-view>
  <FooterArea />
</template>
<script setup lang="ts">
import HeaderArea from '@/views/Main/Layouts/HeaderArea.vue';
import FooterArea from '@/views/Main/Layouts/FooterArea.vue';
</script>

<style scoped></style>
