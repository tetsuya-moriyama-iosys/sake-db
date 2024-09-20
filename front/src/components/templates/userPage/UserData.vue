<template>
  <RadiusImage
    v-if="user.imageBase64"
    :imageSrc="user.imageBase64"
    radius="50px"
  />
  {{ user.name }}
  <div>
    {{ user.profile }}
  </div>
</template>
<script setup lang="ts">
import type { User } from '@/graphQL/User/user';
import RadiusImage from '@/components/parts/common/RadiusImage.vue';

interface Props {
  user: User;
}

const { user } = defineProps<Props>();
</script>
<style scoped></style>
