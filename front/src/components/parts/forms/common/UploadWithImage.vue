<template>
  <img
    v-if="encodedImage"
    :src="`data:image/jpg;base64,${encodedImage}`"
    alt="画像"
  />
  <FormField
    :name="props.name"
    type="file"
    as="input"
    :label="props.label ?? '画像'"
    rules="required|image|size:5000"
    @change="onChange"
  />
  <CommonButton v-if="props.isEnableDelete" @click="onDelete"
    >削除</CommonButton
  >
</template>

<script setup lang="ts">
import FormField from '@/components/parts/forms/core/FormField.vue';
import { ref } from 'vue';
import imageCompression, { type Options } from 'browser-image-compression';
import convertToBase64 from '@/funcs/util/convertToBase64 ';
import CommonButton from '@/components/parts/common/CommonButton.vue';
import { useField } from 'vee-validate';

interface Props {
  name: string;
  default?: string;
  label?: string;
  isEnableDelete?: boolean;
  compressOption?: Options;
}

const defaultCompressOption: Options = {
  maxSizeMB: 1,
  maxWidthOrHeight: 240,
  useWebWorker: true,
};

const props = defineProps<Props>();

// vee-validate用のフィールド定義
const { value: fieldValue } = useField(props.name);

const encodedImage = ref<string | undefined>(props?.default); //base64エンコードしたイメージを格納する
const emit = defineEmits(['onCompressed']); // 親にデータを渡すためのイベント;

const onChange = async (e: Event): Promise<void> => {
  const inputElement = e.target as HTMLInputElement | null;
  if (inputElement?.files == null || inputElement?.files?.length === 0) {
    return;
  }

  try {
    //まずファイルを圧縮する
    const file = inputElement.files[0];
    const compressedFile: File = await imageCompression(
      file,
      props.compressOption ?? defaultCompressOption,
    );
    //圧縮されたファイルをBase64エンコードする
    const encodedStr: string = await convertToBase64(compressedFile);
    encodedImage.value = encodedStr;
    // 圧縮されたBase64データを親コンポーネントにemitで送信
    emit('onCompressed', encodedStr);
  } catch (error) {
    console.error(error);
  }
};

const onDelete = () => {
  encodedImage.value = undefined; // encodedImageをクリア
  fieldValue.value = undefined; // フォームの値をクリア
  // 圧縮されたBase64データを親コンポーネントにemitで送信
  emit('onCompressed', null);
};
</script>

<style scoped></style>
