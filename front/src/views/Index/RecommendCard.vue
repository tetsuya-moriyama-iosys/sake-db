<!--ブックマークレコメンドのカード-->
<template>
  <div v-if="data" class="container">
    <router-link :to="{ name: 'LiquorDetail', params: { id: data.liquor.id } }">
      <div class="top-content">
        <img
          v-if="data.liquor.imageBase64"
          :src="`data:image/jpg;base64,${data.liquor.imageBase64}`"
          :alt="data.liquor.name"
        />
      </div>
    </router-link>
    <div class="middle-content">
      <router-link
        :to="{
          name: 'CategoryNarrowDown',
          params: { id: data.liquor.categoryId },
        }"
        ><p class="category-name">
          {{ data.liquor.categoryName }}
        </p></router-link
      >
      <router-link
        :to="{ name: 'LiquorDetail', params: { id: data.liquor.id } }"
        ><p class="title">{{ data.liquor.name }}</p></router-link
      >
      <hr />
      <div class="comment-container">
        <RadiusImage
          v-if="data.user.imageBase64"
          :imageSrc="data.user.imageBase64"
          radius="5px"
          :alt="data.user.name"
        />
        <router-link :to="{ name: 'UserPage', params: { id: data.user.id } }">
          {{ data.user.name }}
        </router-link>
        <span class="comment"> :{{ data.comment }} </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import RadiusImage from '@/components/parts/common/RadiusImage.vue';
import type { Recommend } from '@/graphQL/Index/recommends';

interface Props {
  data: Recommend;
}

const { data } = defineProps<Props>();
</script>

<style scoped>
div.container {
  display: grid;
  grid-template-rows: 1fr 80px; /* 上側を自動 (1fr)、下側を固定長 (100px) */
  height: 100%; /* コンテナ全体の高さを指定（親要素の高さに依存） */

  border: 1px solid #777;

  div.top-content {
    img {
      width: 100%;
    }
  }
  div.middle-content {
    p.category-name {
      font-size: 75%;
    }
    p.title {
      font-weight: bold;
    }
  }

  div.comment-container {
    span.comment {
      font-size: 10px;
    }
  }
}
</style>
