<template>
  <CommonTag
    :text="tag.text"
    :is-close="!!user"
    @close="isShowCloseDialog = true"
  />
  <YesNoDialog v-model="isShowCloseDialog" :on-yes="deleteTag">
    削除します。宜しいですか？
  </YesNoDialog>
</template>
<script setup lang="ts">
import { ref } from 'vue';

import YesNoDialog from '@/components/parts/common/CommonDialog/Variations/YesNoDialog.vue';
import CommonTag from '@/components/parts/common/CommonTag/CommonTag.vue';
import type { Tag } from '@/graphQL/Liquor/liquor';
import { useUserStore } from '@/stores/userStore';

const { user } = useUserStore();
const { tag } = defineProps<{
  tag: Tag;
}>();

const isShowCloseDialog = ref<boolean>(false);

async function deleteTag() {
  // ここに削除処理を書く
}
</script>

<style scoped></style>
