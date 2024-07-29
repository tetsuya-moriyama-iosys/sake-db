<template>
  <div>
    <button @click="fetcseeddrhDcsata">Get JSON!</button>
    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import axios from 'axios';

export default defineComponent({
  name: 'ExampleComponent',
  setup() {
    const message = ref<string | null>(null);

    const fetchData = async () => {
      try {
        const response = await axios.get(
          'http://localhost:8080/getById/66a1dfb589429cbd8dfefaf1',
        );
        message.value = response.data.message;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    return {
      message,
      fetchData,
    };
  },
});
</script>

<style scoped>
button {
  font-size: 1em;
  padding: 0.5em;
}
</style>
