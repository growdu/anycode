<template>
    <el-form ref="loginForm" :model="loginForm" :rules="loginRules" label-width="80px" class="login-form">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="loginForm.username" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" v-model="loginForm.password" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="login">登录</el-button>
        <el-button type="primary" @click="goToRegisterPage">注册</el-button>
      </el-form-item>
    </el-form>
  </template>
  
  <script>
  export default {
    data() {
      return {
        loginForm: {
          username: '',
          password: ''
        },
        loginRules: {
          username: [{ required: true, message: 'Please input username', trigger: 'blur' }],
          password: [{ required: true, message: 'Please input password', trigger: 'blur' }]
        }
      };
    },
    methods: {
      login() {
        // 登录逻辑
        console.log('Logging in...');
      },
      goToRegisterPage() {
        this.$router.push('/register');
      }
    }
  };
  </script>
  
  <style scoped>
  .login-form {
    width: 300px;
    margin: 0 auto;
  }
  </style>
  