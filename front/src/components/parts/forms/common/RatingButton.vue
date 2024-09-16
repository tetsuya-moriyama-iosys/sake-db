<template>
  <div>
    <div v-if="label">{{ props.label }}</div>
    <div class="rating">
      <span
        v-for="star in Array.from(
          { length: props.maxRating },
          (_, index) => index + 1,
        )"
        :key="star"
        :class="{ active: star <= currentRating }"
        @click="toggleRating(star)"
        class="star"
      >
        {{ star <= currentRating ? '★' : '☆' }}
      </span>
    </div>
    <FormField :name="name" type="hidden" :value="currentRating" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import FormField from '@/components/parts/forms/core/FormField.vue';

interface Props {
  name: string;
  label?: string;
  maxRating?: number; //星の数
}
const props = withDefaults(defineProps<Props>(), {
  maxRating: 5,
});

// 現在の評価（未評価は0）
const currentRating = ref<number>(0);

// 評価の切り替え処理
const toggleRating = (star: number) => {
  // 同じ星をクリックした場合は未評価にリセット
  currentRating.value = currentRating.value === star ? 0 : star;
};
</script>

<style scoped>
.rating {
  display: inline-flex;
  font-size: 1.5rem;
}

.star {
  cursor: pointer;
}

.star.active {
  color: gold;
}
</style>
