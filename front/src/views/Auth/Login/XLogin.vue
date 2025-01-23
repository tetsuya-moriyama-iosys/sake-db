<template>
  <div>
    <button @click="loginWithX">Login with X</button>
  </div>
</template>

<script setup lang="ts">
import axios from 'axios';

import createURL from '@/funcs/util/core/createURL';

async function loginWithX() {
  try {
    // バックエンドからX認証用のURLを取得
    const response = await axios.get(createURL('x/login'));
    const { redirectUrl } = response.data;
    console.log('redirectUrl', redirectUrl);

    // Xの認証ページへリダイレクト
    window.location.href = redirectUrl;
  } catch (error) {
    console.error('Failed to get login URL:', error);
  }
}
</script>

<style scoped></style>
