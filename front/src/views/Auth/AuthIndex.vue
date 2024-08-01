<template>
  <div>
    <h2>Register</h2>
    <form @submit.prevent="registerUser">
      <div>
        <label for="username">Username:</label>
        <input type="text" v-model="username" id="username" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <input type="email" v-model="email" id="email" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" v-model="password" id="password" required />
      </div>
      <button type="submit">Register</button>
    </form>
    <div v-if="error">{{ error }}</div>
    <div v-if="user">User registered successfully: {{ user.username }}</div>
  </div>
</template>

<script lang="ts">
import {
  ApolloClient,
  InMemoryCache,
  gql,
  createHttpLink,
} from '@apollo/client/core';
import { defineComponent, ref } from 'vue';

const httpLink = createHttpLink({
  uri: 'http://localhost:8080/query',
});

const client = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache(),
});

export async function register(
  username: string,
  email: string,
  password: string,
) {
  const REGISTER_MUTATION = gql`
    mutation Register($username: String!, $email: String!, $password: String!) {
      register(username: $username, email: $email, password: $password) {
        id
        username
        email
      }
    }
  `;

  const result = await client.mutate({
    mutation: REGISTER_MUTATION,
    variables: { username, email, password },
  });

  return result.data.register;
}

export default defineComponent({
  name: 'AuthRegister',
  setup() {
    const username = ref<string>('');
    const email = ref<string>('');
    const password = ref<string>('');
    const user = ref<{ id: string; username: string; email: string } | null>(
      null,
    );
    const error = ref<string | null>(null);

    const registerUser = async () => {
      try {
        const registeredUser = await register(
          username.value,
          email.value,
          password.value,
        );
        if (registeredUser) {
          user.value = registeredUser;
          error.value = null;
        } else {
          error.value = 'Registration failed';
        }
      } catch (err) {
        error.value = 'Registration failed';
      }
    };

    return {
      username,
      email,
      password,
      user,
      error,
      registerUser,
    };
  },
});
</script>
