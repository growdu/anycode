<template>
    <el-form ref="registerForm" :model="registerForm" :rules="registerRules" label-width="80px" class="register-form">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="registerForm.username" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" v-model="registerForm.password" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="确认密码" prop="repassword">
        <el-input type="password" v-model="registerForm.password" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="register">注册</el-button>
      </el-form-item>
    </el-form>
  </template>
  
  <script>
  export default {
    data() {
      return {
        registerForm: {
          username: '',
          password: ''
        },
        registerRules: {
          username: [{ required: true, message: 'Please input username', trigger: 'blur' }],
          password: [{ required: true, message: 'Please input password', trigger: 'blur' }]
        }
      };
    },
    methods: {
      register() {
        this.$refs.registerForm.validate(valid => {
          if (valid) {
            // 在这里处理注册逻辑，比如发送请求给后端保存用户信息
            console.log('Registering...');
          } else {
            return false;
          }
        });
      }
    }
  };
  </script>
  
  <style scoped>
  .register-form {
    width: 300px;
    margin: 0 auto;
  }
  </style>
  