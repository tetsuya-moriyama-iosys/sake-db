import { defineStore } from 'pinia';

import type { PostFlavorMap } from '@/graphQL/Liquor/flavorMap';

export const GUEST_STORE = 'guest_flavor_map_store';
const STORAGE_KEY = 'flavor_map_voted';

export type VoteData = PostFlavorMap;

export const guestFlavorMapStore = defineStore(GUEST_STORE, {
  state: () => {
    // 初期化時にローカルストレージからデータを読み込む
    const data: string | null = localStorage.getItem(STORAGE_KEY); //JSON文字列なのでstring[]ではない
    return {
      votedData: (data ? JSON.parse(data) : []) as VoteData[], // 投票済みのID配列
    };
  },

  actions: {
    // 配列に項目を追加し、ローカルストレージに保存
    addItem(item: VoteData) {
      this.votedData.push(item);
      saveToLocalStorage(this.votedData);
    },

    getById(liquorId: string): VoteData | null {
      return this.votedData.find((data) => data.liquorId === liquorId) ?? null;
    },
  },
});

// ローカルストレージに保存
function saveToLocalStorage(data: VoteData[]) {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(data));
}
