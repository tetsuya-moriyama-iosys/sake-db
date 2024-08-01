<template>
  <div>
    <h2>axios新規登録</h2>
    <form @submit.prevent="register">
      <div>
        <label for="name">Name:</label>
        <input v-model="name" type="text" id="name" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <input v-model="email" type="email" id="email" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input v-model="password" type="password" id="password" required />
      </div>
      <button type="submit">Register</button>
    </form>
    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script lang="ts">
import axios from 'axios';

export default {
  data() {
    return {
      name: '' as string,
      email: '' as string,
      password: '' as string,
      message: null as string | null,
    };
  },
  methods: {
    async register() {
      try {
        const response = await axios.post('http://localhost:8080/register', {
          name: this.name,
          email: this.email,
          password: this.password,
        });
        this.message = response.data.message;
      } catch (error) {
        if (axios.isAxiosError(error)) {
          this.message =
            'Registration failed: ' +
            (error.response?.data.error || error.message);
        } else {
          this.message = 'Registration failed: An unknown error occurred';
        }
      }
    },
  },
};
</script>
