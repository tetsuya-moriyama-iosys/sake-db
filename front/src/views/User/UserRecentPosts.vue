<template>
  <div class="title">最近の投稿</div>
  <ul v-if="recentPosts != null && recentPosts.length > 0">
    <li v-for="post in recentPosts" :key="post.liquorId" class="flex">
      <router-link
        :to="{ name: 'LiquorDetail', params: { id: post.liquorId } }"
      >
        <div class="img-container">
          <img
            v-if="post.imageBase64"
            :src="`data:image/jpg;base64,${post.imageBase64}`"
            :alt="post.name"
          />
        </div>
      </router-link>
      <div class="detail flex-1">
        <router-link
          :to="{
            name: 'CategoryNarrowDown',
            params: { id: post.categoryId },
          }"
          ><p class="category-name">{{ post.categoryName }}</p></router-link
        >
        <router-link
          :to="{ name: 'LiquorDetail', params: { id: post.liquorId } }"
          ><p class="title">{{ post.name }}</p></router-link
        >
        <div>
          <DisplayStar :rate="post.rate" />
        </div>
        <div class="comment">
          {{ post.comment }}
        </div>
      </div>
    </li>
  </ul>
  <div v-else>まだ感想の投稿がありません</div>
</template>

<script setup lang="ts">
import DisplayStar from '@/components/parts/common/DisplayStar.vue';
import type { UserLiquor } from '@/graphQL/User/user';

interface Props {
  recentPosts: UserLiquor[] | null;
}

const { recentPosts } = defineProps<Props>();
</script>

<style scoped>
div.title {
  margin-left: 5px;
  border-left: 5px solid black;
  padding-left: 5px;
  font-size: 125%;
  font-weight: bold;
}

li {
  margin-top: 10px;
  margin-bottom: 10px;
  border-bottom: 1px solid #777;

  div.img-container {
    width: 100px;
  }

  img {
    max-height: 100px;
    width: 100px;
  }

  div.detail {
    p.category-name {
      font-size: 75%;
      text-decoration: underline;
    }

    p.title {
      text-decoration: underline;
      font-weight: bold;
    }

    div.comment {
      font-size: 85%;
    }
  }
}
</style>
