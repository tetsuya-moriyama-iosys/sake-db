<template>
  <div v-if="flavorMap">
    <div class="flavor-map-container flex flex-col">
      <div class="text-center order-[-1]">{{ flavorMap.yNames[0] }}</div>
      <div class="text-center order-1">{{ flavorMap.yNames[1] }}</div>
      <div class="flex flex-row">
        <div class="flex items-center order-[-1]">
          {{ flavorMap.xNames[1] }}
        </div>
        <div class="flex items-center order-1">{{ flavorMap.xNames[0] }}</div>
        <div class="relative">
          <!-- 上の矢印 -->
          <div class="absolute top-0 left-1/2 transform -translate-x-1/2">
            <div class="arrow-up"></div>
          </div>

          <!-- 左の矢印 -->
          <div class="absolute top-1/2 left-0 transform -translate-y-1/2">
            <div class="arrow-left"></div>
          </div>

          <!-- 右の矢印 -->
          <div class="absolute top-1/2 right-0 transform -translate-y-1/2">
            <div class="arrow-right"></div>
          </div>

          <!-- 下の矢印 -->
          <div class="absolute bottom-0 left-1/2 transform -translate-x-1/2">
            <div class="arrow-down"></div>
          </div>

          <!-- フレーバーマップのグリッド -->
          <div class="flavor-map-grid-container">
            <div
              v-for="(_, yIndex) in 21"
              :key="10 - yIndex"
              class="grid-row flex"
            >
              <FlavorCell
                v-for="(_, xIndex) in 21"
                :key="xIndex - 10"
                :cellData="
                  flavorMap.mapData.find(
                    (data: FlavorCellType) =>
                      data.x === xIndex - 10 && data.y === 10 - yIndex,
                  )!
                "
                class="grid-cell"
                :isSelectable="isLogin || votedCoordinates == null"
                @click="
                  () => {
                    savedCoordinates = { x: xIndex - 10, y: 10 - yIndex };
                    if (isLogin || votedCoordinates == null) {
                      //なんか到達不能って出てるけど、そんなことない......
                      isShowPostDialog = true;
                    }
                  }
                "
                :isSelected="
                  votedCoordinates?.x === xIndex - 10 &&
                  votedCoordinates?.y === 10 - yIndex
                "
              />
            </div>
          </div>
        </div>
      </div>
    </div>
    <div>
      <p>
        投票数：{{ flavorMap.guestFullAmount + flavorMap.userFullAmount }}人
      </p>
      <p>登録済ユーザ：{{ flavorMap.userFullAmount }}人</p>
      <p>未登録ユーザ：{{ flavorMap.guestFullAmount }}人</p>
    </div>
  </div>
  <YesNoDialog
    v-if="savedCoordinates && flavorMap"
    v-model="isShowPostDialog"
    :on-yes="onSubmit"
  >
    <div v-if="savedCoordinates.x >= 0">
      {{ flavorMap.xNames[0] }}: {{ savedCoordinates.x }}
    </div>
    <div v-else>{{ flavorMap.xNames[1] }}: {{ savedCoordinates.x * -1 }}</div>
    <div v-if="savedCoordinates.y >= 0">
      {{ flavorMap.yNames[0] }}: {{ savedCoordinates.y }}
    </div>
    <div v-else>{{ flavorMap.yNames[1] }}: {{ savedCoordinates.y * -1 }}</div>
    <div>
      登録しますか？
      <div v-if="!isLogin" class="text-red-800">
        ※(注意)未ログイン状態の場合、一度投票したら変えられません。
      </div>
    </div>
  </YesNoDialog>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

import FlavorCell from '@/components/blocks/FlavorMap/FlavorCell.vue';
import YesNoDialog from '@/components/parts/common/CommonDialog/Variations/YesNoDialog.vue';
import { useMutation, useQuery } from '@/funcs/composable/useQuery/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import {
  type Coordinates,
  type FlavorCell as FlavorCellType,
  type FlavorMap,
  type FlavorMapResponse,
  GetFlavorMap,
  GetVoted,
  PostFlavorMap,
  type VotedResponse,
} from '@/graphQL/Liquor/flavorMap';
import type { Liquor } from '@/graphQL/Liquor/liquor';
import { guestFlavorMapStore } from '@/stores/guestFlavorMapStore';
import { useUserStore } from '@/stores/userStore/userStore';

interface Props {
  liquor: Liquor;
}

const { isLogin } = useUserStore();
const { getById, addItem } = guestFlavorMapStore();
const { execute: fetch } = useMutation<FlavorMapResponse>(GetFlavorMap, {
  isAuth: true,
});
const { fetch: votedFetch } = useQuery<VotedResponse>(GetVoted, {
  isAuth: true,
});
const { execute: post } = useMutation<FlavorMapResponse>(PostFlavorMap, {
  isAuth: true,
});
const toast = useToast();

const { liquor } = defineProps<Props>();
const flavorMap = ref<FlavorMap | null>(null);
const isShowPostDialog = ref<boolean>(false);
const votedCoordinates = ref<Coordinates | null>(null);
const savedCoordinates = ref<Coordinates | null>(null);

onMounted(() => {
  void onFetch();
});

async function onFetch() {
  void setVoted(); //再読込ごとに投票済みかどうかを確認
  const response: FlavorMapResponse = await fetch(
    { liquorId: liquor.id },
    {
      fetchPolicy: 'no-cache',
    },
  );
  flavorMap.value = response.getFlavorMap;
}

async function setVoted(): Promise<void> {
  if (isLogin) {
    const response: VotedResponse = await votedFetch(
      { liquorId: liquor.id },
      {
        fetchPolicy: 'no-cache',
      },
    );
    votedCoordinates.value = response.getVoted;
    return;
  }
  //ゲストユーザーはローカルストレージから取得
  votedCoordinates.value = getById(liquor.id);
}

async function onSubmit() {
  await post({
    input: {
      liquorId: liquor.id,
      x: savedCoordinates.value!.x,
      y: savedCoordinates.value!.y,
    },
  });
  toast.showToast({
    message: '投票が完了しました',
  });
  if (!isLogin) {
    addItem({
      liquorId: liquor.id,
      x: savedCoordinates.value!.x,
      y: savedCoordinates.value!.y,
    });
  }
  void onFetch();
}
</script>

<style scoped lang="scss">
div.flavor-map-container {
  width: calc(420px + 2em);
  height: calc(420px + 2em);
}

div.flavor-map-grid-container {
  background-color: #eee;
  div.grid-row {
    height: calc(420px / 21);
  }
}

// 矢印のスタイル
.arrow-up,
.arrow-down,
.arrow-left,
.arrow-right {
  width: 0;
  height: 0;
  border-style: solid;
}

// 矢印の具体的なスタイル設定
.arrow-up {
  border-width: 0 10px 10px 10px;
  border-color: transparent transparent #333 transparent;
}

.arrow-down {
  border-width: 10px 10px 0 10px;
  border-color: #333 transparent transparent transparent;
}

.arrow-left {
  border-width: 10px 10px 10px 0;
  border-color: transparent #333 transparent transparent;
}

.arrow-right {
  border-width: 10px 0 10px 10px;
  border-color: transparent transparent transparent #333;
}
</style>
