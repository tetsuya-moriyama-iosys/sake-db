<!--カテゴリ選択-->

<template>
  <div>{{ label }}</div>
  <div>
    <div v-for="(level, index) in levels" :key="index">
      <select
        v-model="selectedValues[index]"
        @change="handleChange(index)"
        @blur="
          () => {
            void validate();
          }
        "
        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-1.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
      >
        <option :value="undefined">未選択</option>
        <option
          v-for="category in level"
          :key="category.id"
          :value="category.id"
        >
          {{ category.name }}
        </option>
      </select>
    </div>
    <FormField as="input" type="hidden" :value="finalCategoryId" :name="name" />
  </div>
  <div v-if="errorMessage" class="error">
    <ErrorMessage :name="name" />
  </div>
</template>

<script setup lang="ts">
import { ErrorMessage, useField } from 'vee-validate';
import { computed, type ComputedRef, ref, watch } from 'vue';

import FormField from '@/components/parts/forms/core/FormField.vue';
import useQuery from '@/funcs/composable/useQuery';
import {
  type Categories,
  type Category,
  GET_QUERY,
} from '@/graphQL/Category/categories';

//propsのセット
const props = defineProps<{
  label?: string;
  name: string;
  initialId?: number | null;
}>();

//階層構造の記憶
const selectedValues = ref<Array<string | undefined>>([]);
//現時点で表示しているデータ
const levels = ref<Category[][]>([]);

const { fetch } = useQuery<Categories>(GET_QUERY);

// vee-validate用のフィールド定義
const { value: hiddenField, errorMessage, validate } = useField(props.name);

// 初期IDが変更されたら、選択肢を再初期化
watch(
  () => props.initialId,
  async (newVal: number | null | undefined) => {
    const { categories: response } = await fetch({
      fetchPolicy: 'no-cache', //カテゴリが途中で変更されると、意図しない変更になるリスクがあるので再取得
    });
    levels.value = [response]; // 最初の階層を設定
    initializeSelections(newVal, response); // 初期値で選択肢を初期化
    hiddenField.value = newVal?.toString(); // hiddenFieldにも設定
    console.log('hiddenField.value:', hiddenField.value);
    //idが空値の場合、セレクトボックスの初期化が必要
    if (props.initialId == null) {
      selectedValues.value = [];
    }
  },
  { immediate: true }, // 初回に監視対象がある場合も実行
);

//変更時の操作
const handleChange = (index: number) => {
  // 選択が変更されたらそこから先の選択肢を一旦削除(indexが変更された階層)
  selectedValues.value = selectedValues.value.slice(0, index + 1);
  levels.value = levels.value.slice(0, index + 1);

  if (selectedValues.value[index] === undefined) {
    return;
  }

  // 新しい選択肢を追加
  const selectedId = selectedValues.value[index];
  const selectedCategory = findCategoryById(
    levels.value[0], //大元のカテゴリを起点にして検索開始
    parseInt(selectedId as string), //ここに到達する時点でundefinedではない
  );

  //見つかったカテゴリが子カテゴリを持っていれば、選択肢に追加
  if (selectedCategory && selectedCategory.children) {
    levels.value.push(selectedCategory.children);
  }
};

//指定したidを持つデータを再帰的に検索し、取得する
const findCategoryById = (
  categories: Category[],
  id: number,
): Category | undefined => {
  for (const category of categories) {
    if (category.id === id) return category;
    if (category.children) {
      const found = findCategoryById(category.children, id);
      if (found) return found;
    }
  }
};

//カテゴリリストの一番子のIDが、現時点で取得されている(最後に選択された)もの
const finalCategoryId: ComputedRef<string> = computed(() => {
  const lastSelected: string | undefined =
    selectedValues.value[selectedValues.value.length - 1];
  if (lastSelected !== undefined) {
    return lastSelected as string;
  }
  //未選択の場合
  if (selectedValues.value.length === 1) {
    //一番親のカテゴリだった場合は空値を返す
    return '';
  }
  //それ以外の場合は、更にその親の値を返せばOK
  return selectedValues.value[selectedValues.value.length - 2] as string;
});

// finalCategoryIdの変更を検知してcategoryIdをセットする
watch(finalCategoryId, (newVal) => {
  hiddenField.value = newVal;
});

// 指定された値に基づいてセレクトボックスを設定する関数
const initializeSelections = (
  id: number | null | undefined, //未設定か明示的にnullが渡されてくるかでnullかundefined
  categories: Category[],
) => {
  //該当するIDまでのカテゴリ配列を取得
  const path: Category[] = findCategoryPathById(categories, id);
  if (path.length > 0) {
    path.forEach((category, index) => {
      selectedValues.value[index] = category.id.toString();
      if (category.children) {
        //見つかった子配列を上書き
        levels.value[index + 1] = category.children;
      }
    });
  }
};

// カテゴリIDから親カテゴリまでのパスを見つける関数
const findCategoryPathById = (
  categories: Category[],
  id: number | null | undefined,
  path: Category[] = [],
): Category[] => {
  for (const category of categories) {
    if (category.id === id) {
      return [...path, category];
    }
    if (category.children) {
      const result = findCategoryPathById(category.children, id, [
        ...path,
        category,
      ]);
      if (result.length > 0) {
        return result;
      }
    }
  }
  return [];
};
</script>

<style scoped>
div.error {
  color: red;
  font-size: 75%;
}
</style>
