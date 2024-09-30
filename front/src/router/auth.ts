import RegisterPage from '@/views/Auth/RegisterPage.vue';
import LoginPage from '@/views/Auth/LoginPage.vue';
import PasswordReset from '@/views/Auth/PasswordReset.vue';
import PasswordResetExe from '@/views/Auth/PasswordResetExe.vue';

const AuthRouter = {
  path: 'auth',
  children: [
    //要認証ページはmypage以下にある。あくまでも「認証関係」ということでauthという命名にしている。
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
    {
      path: 'password',
      children: [
        {
          path: 'reset',
          name: 'PasswordReset',
          component: PasswordReset,
        },
        {
          path: 'reset/:token',
          component: PasswordResetExe,
        },
      ],
    },
  ],
};

export default AuthRouter;
