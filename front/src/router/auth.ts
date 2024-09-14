import RegisterPage from '@/views/Auth/RegisterPage.vue';
import LoginPage from '@/views/Auth/LoginPage.vue';

const AuthRouter = {
  path: 'auth',
  children: [
    {
      path: 'register',
      name: 'Register',
      component: RegisterPage,
    },
    {
      path: 'login',
      name: 'Login',
      component: LoginPage,
    },
  ],
};

export default AuthRouter;
