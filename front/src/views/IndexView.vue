<template>
  <div>
    <button @click="fetchMessages">Fetch Messages</button>
    <div v-if="loading">Loading...</div>
    <div v-if="error">{{ error.message }}</div>
    <ul v-if="data">
      <li v-for="message in data.messages" :key="message.id">
        {{ message.message }}
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { ref, type Ref } from 'vue';
import { ApolloError } from '@apollo/client/core';
import client from '@/apolloClient';

// GraphQLクエリの型定義
import { GET_MESSAGES } from '@/graphQL/query/query';
import { type MessagesData, type Message } from '@/graphQL/type/type';

export default {
  setup() {
    const loading: Ref<boolean> = ref(false);
    const error: Ref<ApolloError | null> = ref(null);
    const data: Ref<MessagesData | null> = ref(null);

    const fetchMessages = async (): Promise<void> => {
      loading.value = true;
      error.value = null;

      try {
        const result = await client.query<{ messages: Message[] }>({
          query: GET_MESSAGES,
        });
        data.value = result.data;
      } catch (err) {
        if (err instanceof ApolloError) {
          error.value = err;
        } else {
          console.error(err);
        }
      } finally {
        loading.value = false;
      }
    };

    return {
      loading,
      error,
      data,
      fetchMessages,
    };
  },
};
</script>
