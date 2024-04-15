import { createRouter, createWebHistory } from 'vue-router';
import Login from '../view/Login.vue';
import Register from '../view/Register.vue';
import Home from '../view/Home.vue';
import ChangePwd from '../view/ChangePwd.vue';

const routes = [
  {
    path: '/',
    redirect: '/login' // 将根路径重定向到登录页面
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/login/:username',
    name: 'LoginWithParam',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/home:username',
    name: 'Home',
    component: Home,
    props: true
  },
  {
    path: '/change:username',
    name: 'Change',
    component: ChangePwd,
    props: true
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
