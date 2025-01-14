<template>
  <CommonTag
    :text="tag.text"
    :is-close="!!user"
    @close="isShowDeleteDialog = true"
  />
  <YesNoDialog v-model="isShowDeleteDialog" :on-yes="deleteTag">
    削除します。宜しいですか？
  </YesNoDialog>
</template>
<script setup lang="ts">
import { ref } from 'vue';

import YesNoDialog from '@/components/parts/common/CommonDialog/Variations/YesNoDialog.vue';
import CommonTag from '@/components/parts/common/CommonTag/CommonTag.vue';
import { useMutation } from '@/funcs/composable/useQuery/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import type { Tag } from '@/graphQL/Liquor/liquor';
import { DeleteTag } from '@/graphQL/Liquor/tags';
import { useUserStore } from '@/stores/userStore/userStore';
const { user } = useUserStore();
const toast = useToast();
const { execute } = useMutation(DeleteTag, {
  isAuth: true,
});

const emit = defineEmits<{
  delete: [string];
}>();

const { tag } = defineProps<{
  tag: Tag;
}>();

const isShowDeleteDialog = ref<boolean>(false);

async function deleteTag() {
  // ここに削除処理を書く
  await execute({ id: tag.id });
  toast.showToast({
    message: 'タグの削除に成功しました',
  });
  emit('delete', tag.id);
}
</script>

<style scoped></style>
