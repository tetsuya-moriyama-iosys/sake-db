import { defineStore } from 'pinia';

export const useSelectedCategoryStore = defineStore({
  id: 'selectedCategory',
  state: () => ({
    content: null as number | null, // 型を明示的に指定
  }),
  actions: {
    updateContent(newId: number | null) {
      this.content = newId; // アクションでstateの更新を行う
    },
  },
});
