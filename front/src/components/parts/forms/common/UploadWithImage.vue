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
</template>

<script setup lang="ts">
import FormField from '@/components/parts/forms/core/FormField.vue';
import { ref } from 'vue';
import imageCompression from 'browser-image-compression';
import convertToBase64 from '@/funcs/util/convertToBase64 ';

interface Props {
  name: string;
  default?: string;
  label?: string;
}

const props = defineProps<Props>();

const encodedImage = ref<string | undefined>(props?.default); //base64エンコードしたイメージを格納する
const emit = defineEmits(['onCompressed']); // 親にデータを渡すためのイベント;

const onChange = async (e: Event): Promise<void> => {
  const inputElement = e.target as HTMLInputElement | null;
  if (
    inputElement == null ||
    inputElement?.files === null ||
    inputElement?.files?.length === 0
  ) {
    return;
  }

  const file = inputElement.files[0];
  console.log(file);
  if (file == null) return;
  const options = {
    maxSizeMB: 1,
    maxWidthOrHeight: 1920,
    useWebWorker: true,
  };
  try {
    const compressedFile = await imageCompression(file, options);
    console.log(
      'compressedFile instanceof Blob',
      compressedFile instanceof Blob,
    ); // true
    console.log(`compressedFile size ${compressedFile.size / 1024 / 1024} MB`); // smaller than maxSizeMB
    encodedImage.value = await convertToBase64(compressedFile);
    // 圧縮されたBase64データを親コンポーネントにemitで送信
    emit('onCompressed', compressedFile);

    //await uploadToServer(compressedFile); // write your own logic
  } catch (error) {
    console.error(error);
  }
};
</script>

<style scoped></style>
