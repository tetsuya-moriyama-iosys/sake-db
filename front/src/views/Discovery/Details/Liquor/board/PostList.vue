<template>
  <div>みんなの感想</div>
  <div v-for="(post, index) in props.posts" :key="index" class="flex">
    <div>
      <router-link
        v-if="post.userId"
        :to="{ name: 'UserPage', params: { id: post.userId } }"
        >{{ post.userName }}</router-link
      >
      <span v-else>名無し</span>
      :
    </div>
    <div>
      <div v-if="post.rate">
        (<span v-for="i in 5" :key="i">
          <span v-if="i <= post.rate">★</span>
          <span v-else>☆</span> </span
        >)
      </div>
      <div v-else>(未評価)</div>
    </div>
    <div>{{ post.text }}</div>
  </div>
</template>
<script setup lang="ts">
import type { Post } from '@/graphQL/Liquor/board';

interface Props {
  posts: Post[];
}

const props = defineProps<Props>();
</script>

<style scoped></style>
