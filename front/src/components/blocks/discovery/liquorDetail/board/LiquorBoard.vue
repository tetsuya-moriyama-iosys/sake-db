<template>
  <PostForm :liquor-id="props.liquorId" />
  <PostList />
</template>

<script setup lang="ts">
import PostForm from '@/components/blocks/discovery/liquorDetail/board/PostForm.vue';
import PostList from '@/components/blocks/discovery/liquorDetail/board/PostList.vue';

interface Props {
  liquorId: string;
}

const props = defineProps<Props>();
</script>

<style scoped></style>
