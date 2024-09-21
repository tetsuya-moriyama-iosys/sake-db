<template>
  <RadiusImage
    v-if="userDetail.user.imageBase64"
    :imageSrc="userDetail.user.imageBase64"
    radius="50px"
  />
  {{ userDetail.user.name }}
  <div>
    {{ userDetail.user.profile }}
  </div>
</template>
<script setup lang="ts">
import type { UserDetail } from '@/graphQL/User/user';
import RadiusImage from '@/components/parts/common/RadiusImage.vue';

interface Props {
  userDetail: UserDetail;
}

const { userDetail } = defineProps<Props>();
</script>
<style scoped></style>
