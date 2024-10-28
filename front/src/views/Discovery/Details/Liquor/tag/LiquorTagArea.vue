<template>
  <div>
    タグ一覧
    <div v-if="tags.length > 0">
      <LiquorTag
        v-for="tag in tags"
        :tag="tag"
        :key="tag.id"
        @delete="deleted"
      />
    </div>
    <div v-else>登録されたタグはありません</div>
    <TagInput v-if="user" :liquor-id="props.liquorId" @submitted="submitted" />
  </div>
</template>
<script setup lang="ts">
import { onMounted, ref } from 'vue';

import useQuery from '@/funcs/composable/useQuery';
import type { Tag } from '@/graphQL/Liquor/liquor';
import { FetchTags, type GetTagsResponse } from '@/graphQL/Liquor/tags';
import { useUserStore } from '@/stores/userStore';
import LiquorTag from '@/views/Discovery/Details/Liquor/tag/LiquorTag.vue';
import TagInput from '@/views/Discovery/Details/Liquor/tag/TagInput.vue';

const { user } = useUserStore();
const { fetch } = useQuery<GetTagsResponse>(FetchTags);

const props = defineProps<{
  liquorId: string;
}>();

const tags = ref<Tag[]>([]);

onMounted(async () => {
  const response: GetTagsResponse = await fetch({ liquorId: props.liquorId });
  tags.value = response.getTags;
});

//新しいタグが投稿されたら、画面上に反映する
function submitted(newTag: Tag) {
  tags.value = [...tags.value, newTag];
}

//削除処理が成功したら、画面上から削除する
function deleted(deleteId: string) {
  tags.value = tags.value.filter((tag) => tag.id !== deleteId);
}
</script>

<style scoped></style>
