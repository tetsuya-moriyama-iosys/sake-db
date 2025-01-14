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
  <div>
    <BookmarkedList :id="userDetail.user.id" />
  </div>
  <div class="user-recent-posts mt-5">
    <UserRecentPosts :recent-posts="userDetail.evaluateList.recentComments" />
  </div>
</template>

<script setup lang="ts">
import RadiusImage from '@/components/parts/common/RadiusImage.vue';
import type { UserDetail } from '@/graphQL/User/user';
import { useUserStore } from '@/stores/userStore/userStore';
import BookMarkButton from '@/views/User/BookMarkButton.vue';
import BookmarkedList from '@/views/User/BookmarkedList.vue';
import UserPosts from '@/views/User/UserPosts.vue';
import UserRecentPosts from '@/views/User/UserRecentPosts.vue';

interface Props {
  userDetail: UserDetail;
}
const userStore = useUserStore();
const { userDetail } = defineProps<Props>();
</script>

<style scoped></style>
