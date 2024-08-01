import AuthRegister from '@/views/Auth/AuthIndex.vue';
import NestComponent from '@/views/NestComponent.vue';

const AuthRouter = {
  path: '/auth',
  name: 'Auth',
  component: NestComponent, // 親ルート用の空コンポーネント
  children: [
    {
      path: 'register',
      name: 'AuthRegister',
      component: AuthRegister,
    },
  ],
};

export default AuthRouter;
