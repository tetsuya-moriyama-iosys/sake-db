import { defineStore } from 'pinia';

export const SIDEBAR_STORE = 'selectedCategory';
export const useSelectedCategoryStore = defineStore({
  id: SIDEBAR_STORE,
  state: () => ({
    content: null as number | null, // 型を明示的に指定
    isReloadFlg: false, //左サイドメニューを強制的にリロードするためのフラグ
  }),
  actions: {
    updateContent(newId: number | null) {
      this.content = newId; // アクションでstateの更新を行う
    },
    reload() {
      this.isReloadFlg = true;
    },
    setReloadFlgFalse() {
      this.isReloadFlg = false;
    },
  },
});
