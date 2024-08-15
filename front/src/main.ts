import './styles/main.css';

import { createApp, h, provide } from 'vue';
import { createPinia } from 'pinia';
import './styles/tailwind.css';

import App from './App.vue';
import router from './router';
import client from '@/apolloClient';
import { DefaultApolloClient } from '@vue/apollo-composable';

const app = createApp({
  setup() {
    provide(DefaultApolloClient, client);
  },
  render: () => h(App),
});

app.use(createPinia());
app.use(router);

app.mount('#app');
