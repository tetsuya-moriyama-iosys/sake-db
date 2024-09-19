<template>
  <router-link :to="{ name: 'MyPageEdit' }">
    <div v-if="userStore.user" class="flex">
      <RadiusImage
        v-if="userStore.user?.imageBase64"
        :imageSrc="userStore.user.imageBase64"
        radius="30px"
      />
      <div>
        {{ userStore.user.name }}
      </div>
    </div>
  </router-link>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/userStore';
import RadiusImage from '@/components/parts/common/RadiusImage.vue';

const userStore = useUserStore();

// watch(
//   () => userStore.user,
//   (newUser) => {},
// );
</script>

<style scoped></style>
