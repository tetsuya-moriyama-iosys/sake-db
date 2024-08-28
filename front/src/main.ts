import './styles/main.css';

import { createApp, h, provide } from 'vue';
import './styles/tailwind.css';

import App from './App.vue';
import client from '@/apolloClient';
import { DefaultApolloClient } from '@vue/apollo-composable';
import { registerPlugins } from '@/plugins';

const app = createApp({
  setup() {
    provide(DefaultApolloClient, client);
  },
  render: () => h(App),
});

// プラグインの登録
registerPlugins(app);

app.mount('#app');
