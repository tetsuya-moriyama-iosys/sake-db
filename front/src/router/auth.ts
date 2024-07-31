import Index from '../views/IndexView.vue';
import NestComponent from '../views/NestComponent.vue';

const AuthRouter = {
  path: '/auth',
  name: 'Auth',
  component: NestComponent, // 親ルート用の空コンポーネント
  children: [
    {
      path: 'team',
      name: 'AboutTeam',
      component: Index,
    },
  ],
};

export default AuthRouter;
