<template>
  <div class="sm:hidden">
    <select
      id="tabs"
      class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
    >
      <option v-for="rate in [5, 4, 3, 2, 1]" :key="`smart_rate_${rate}`">
        <DisplayStar :rate="rate" />
      </option>
      <option>未評価</option>
    </select>
  </div>
  <ul
    class="hidden text-sm font-medium text-center text-gray-500 rounded-lg shadow sm:flex dark:divide-gray-700 dark:text-gray-400"
  >
    <li
      v-for="rate in [5, 4, 3, 2, 1, null]"
      class="w-full focus-within:z-10"
      :key="`rate_${rate}`"
      @click="onRateSelect(rate)"
    >
      <span
        :class="
          selectedRate === rate
            ? 'text-gray-900 bg-gray-100 rounded-s-lg active dark:bg-gray-700 dark:text-white'
            : 'bg-white dark:bg-gray-800 hover:text-gray-700 hover:bg-gray-50 dark:hover:text-white dark:hover:bg-gray-700'
        "
        class="inline-block w-full p-4 border-r border-gray-200 dark:border-gray-700 focus:ring-4 focus:ring-blue-300 focus:outline-none"
        :aria-current="selectedRate === rate ? 'page' : undefined"
        ><span v-if="rate != null"><DisplayStar :rate="rate" /></span>
        <span v-else>未評価</span></span
      >
    </li>
  </ul>
  <div>
    <RatedLiquorList :liquor-list="liquorList" />
  </div>
</template>

<script setup lang="ts">
import type { EvaluateList, UserLiquor } from '@/graphQL/User/user';
import { ref } from 'vue';
import RatedLiquorList from '@/components/blocks/userPage/RatedLiquorList.vue';
import DisplayStar from '@/components/parts/common/DisplayStar.vue';

interface Props {
  evaluates: EvaluateList;
}

const { evaluates } = defineProps<Props>();

const selectedRate = ref<number | null>(5);
const liquorList = ref<UserLiquor[] | null>(evaluates.rate5Liquors);

const rate5List: UserLiquor[] | null = evaluates.rate5Liquors;
const rate4List: UserLiquor[] | null = evaluates.rate4Liquors;
const rate3List: UserLiquor[] | null = evaluates.rate3Liquors;
const rate2List: UserLiquor[] | null = evaluates.rate2Liquors;
const rate1List: UserLiquor[] | null = evaluates.rate1Liquors;
const noRateList: UserLiquor[] | null = evaluates.noRateLiquors;

function onRateSelect(rate: number | null): void {
  selectedRate.value = rate;
  switch (rate) {
    case 5:
      liquorList.value = rate5List;
      break;
    case 4:
      liquorList.value = rate4List;
      break;
    case 3:
      liquorList.value = rate3List;
      break;
    case 2:
      liquorList.value = rate2List;
      break;
    case 1:
      liquorList.value = rate1List;
      break;
    default:
      liquorList.value = noRateList;
  }
}
</script>

<style scoped></style>
