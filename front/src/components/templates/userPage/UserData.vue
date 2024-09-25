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
  <BookMarkButton
    v-if="userStore.isLogin && userDetail.user.id != userStore.user?.id"
    :target-id="userDetail.user.id"
  />
  <div>
    <UserPosts :evaluates="userDetail.evaluateList" />
  </div>
  <div class="user-recent-posts mt-5">
    <UserRecentPosts :recent-posts="userDetail.evaluateList.recentComments" />
  </div>
</template>

<script setup lang="ts">
import type { UserDetail } from '@/graphQL/User/user';
import RadiusImage from '@/components/parts/common/RadiusImage.vue';
import UserPosts from '@/components/templates/userPage/UserPosts.vue';
import UserRecentPosts from '@/components/templates/userPage/UserRecentPosts.vue';
import BookMarkButton from '@/components/blocks/userPage/BookMarkButton.vue';
import { useUserStore } from '@/stores/userStore';

interface Props {
  userDetail: UserDetail;
}
const userStore = useUserStore();
const { userDetail } = defineProps<Props>();
</script>
<style scoped></style>
