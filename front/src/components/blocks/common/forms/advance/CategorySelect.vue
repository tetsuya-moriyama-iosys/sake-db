<!--カテゴリ選択-->

<template>
  <div>
    {{ label }}
  </div>
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
import { computed, onMounted, ref, watch } from 'vue';
import { ErrorMessage, useField } from 'vee-validate';
import { type Category } from '@/type/common/liquor/Category';
import useQuery from '@/funcs/composable/useQuery';
import {
  type Categories,
  GET_QUERY,
} from '@/type/api/common/liquor/Categories';
import FormField from '@/components/parts/forms/core/FormField.vue';

//propsのセット
const {
  label = 'カテゴリ',
  name,
  initialId,
} = defineProps<{
  label?: string;
  name: string;
  initialId?: number | null;
}>();

//階層構造の記憶
const selectedValues = ref<string[]>([]);
//現時点で表示しているデータ
const levels = ref<Category[][]>([]);

const { fetch } = useQuery<Categories>(GET_QUERY);

// vee-validate用のフィールド定義
const { value: hiddenField, errorMessage, validate } = useField(name);

//変更時の操作
const handleChange = (index: number) => {
  // 選択が変更されたらそこから先の選択肢を一旦削除(indexが変更された階層)
  selectedValues.value = selectedValues.value.slice(0, index + 1);
  levels.value = levels.value.slice(0, index + 1);

  // 新しい選択肢を追加
  const selectedId = selectedValues.value[index];
  const selectedCategory = findCategoryById(
    levels.value[0], //大元のカテゴリを起点にして検索開始
    parseInt(selectedId),
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
const finalCategoryId = computed(() => {
  const lastSelected = selectedValues.value[selectedValues.value.length - 1];
  return lastSelected || '';
});

// finalCategoryIdの変更を検知してcategoryIdをセットする
watch(finalCategoryId, (newVal) => {
  hiddenField.value = newVal;
});

// 読み込み時にカテゴリ情報をAPIから取得
onMounted(async () => {
  const { categories: response } = await fetch();
  levels.value = [response]; // 最初の階層を設定

  if (initialId != null) {
    initializeSelections(initialId, response);
    hiddenField.value = initialId.toString(); // 初期値をhiddenFieldに設定
  }
});

// 指定された値に基づいてセレクトボックスを設定する関数
const initializeSelections = (id: number, categories: Category[]) => {
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
  id: number,
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
