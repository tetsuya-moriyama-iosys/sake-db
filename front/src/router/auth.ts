import { defineComponent, h } from 'vue';
import AuthRegister from '@/views/Auth/AuthIndex.vue';

const AuthRouter = {
  path: '/auth',
  name: 'Auth',
  component: defineComponent({
    render() {
      return h('router-view'); //ダミーコンポーネント
    },
  }),
  children: [
    {
      path: 'register',
      name: 'AuthRegister',
      component: AuthRegister,
    },
  ],
};

export default AuthRouter;
