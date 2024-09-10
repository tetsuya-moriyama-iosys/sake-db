import AuthIndex from '@/views/Auth/AuthIndex.vue';

const AuthRouter = {
  path: 'auth',
  children: [
    {
      path: 'register',
      name: 'Register',
      component: AuthIndex,
    },
    {
      path: 'login',
      name: 'Login',
      component: AuthIndex,
    },
  ],
};

export default AuthRouter;
