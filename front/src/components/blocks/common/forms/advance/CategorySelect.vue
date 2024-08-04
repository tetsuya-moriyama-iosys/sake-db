<!--カテゴリ選択-->

<template>
  <!--ID値を格納する実態-->
  <FormField :name="name" type="hidden" />
  <MasterCategorySelect v-bind="propsObj" />
</template>

<script setup lang="ts">
import MasterCategorySelect from '@/components/blocks/common/forms/MasterCategorySelect.vue';
import type { SelectFieldProps } from '@/components/parts/forms/core/SelectField.vue';
import FormField from '@/components/parts/forms/core/FormField.vue';
import { gql } from '@apollo/client/core';
import { onMounted, ref } from 'vue';
import client from '@/apolloClient';

const GET_QUERY = gql`
  query {
    categories {
      id
      name
      children {
        id
        name
        children {
          id
          name
        }
      }
    }
  }
`;

const {
  label = 'カテゴリ',
  name,
  ...props
} = defineProps<Omit<SelectFieldProps, 'options'>>();

const loading = ref<boolean>(false);
const error = ref<unknown>(null);
const data = ref(null);
const fetch = async () => {
  loading.value = true;
  error.value = null;
  try {
    const result = await client.query({
      query: GET_QUERY,
    });
    console.log('result:', result);
    data.value = result.data;
  } catch (err) {
    error.value = err;
  } finally {
    loading.value = false;
  }
};

// onMounted フックを使用して fetchMessages 関数を呼び出します。
onMounted(async () => {
  await fetch();
  console.log('data:', data);
});

//オブジェクトを作らないとバインディングされないので詰め替える
const propsObj = { ...props, label, name: 'master-category' }; //name値はフォームとして使われないので固定値
</script>

<style scoped></style>
