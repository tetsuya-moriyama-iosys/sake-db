<template>
  <router-link :to="{ name: 'MyPageIndex' }">
    <div v-if="userStore.user" class="flex container">
      <RadiusImage
        v-if="userStore.user.imageBase64"
        :imageSrc="userStore.user.imageBase64"
        radius="10px"
      />
      <div>
        {{ userStore.user.name }}
      </div>
    </div>
  </router-link>
</template>

<script setup lang="ts">
import RadiusImage from '@/components/parts/common/RadiusImage.vue';
import { useUserStore } from '@/stores/userStore/userStore';

const userStore = useUserStore();
</script>

<style scoped></style>
