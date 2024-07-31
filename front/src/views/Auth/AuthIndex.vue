<template>
  <div>
    <h1>Hello World</h1>
    <ul>
      <li v-for="item in helloWorld" :key="item._id">{{ item.message }}</li>
    </ul>
  </div>
</template>

<script>
import { useQuery, gql } from '@vue/apollo-composable';

const GET_HELLO_WORLD = gql`
  query GetHelloWorld {
    helloWorld {
      _id
      message
    }
  }
`;

export default {
  setup() {
    const { result, loading, error } = useQuery(GET_HELLO_WORLD);

    return {
      helloWorld: result,
      loading,
      error,
    };
  },
};
</script>
