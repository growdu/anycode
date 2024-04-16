<template>
  <div class="login-container">
    <h1 class="login-title">AnyCode</h1>
    <el-form ref="loginForm" :model="loginForm" :rules="loginRules" label-width="80px" label-position="top" class="login-form">
      <el-form-item label="用户名" prop="username" style="align-items: flex-start;">
        <el-input v-model="loginForm.username" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password" style="align-items: flex-start;">
        <el-input type="password" v-model="loginForm.password" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="login" style="width: 100%">登录</el-button>
      </el-form-item>
    </el-form>
    <div v-if="errorMessage" class="error-message">{{ errorMessage }}</div>
  </div>
</template>

<script>
import { ElMessageBox } from 'element-plus';

export default {
  data() {
    return {
      loginForm: {
        username: '',
        password: ''
      },
      loginRules: {
        username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
      },
      errorMessage: '' // 用于显示错误信息
    };
  },
  created() {
    this.getRouterData();
  },
  methods: {
    async login() {
      // 在这里获取用户名，并通过路由参数传递给 home 界面
      try {
        const response = await this.$axios.post('/login', {
          username: this.loginForm.username,
          password: this.loginForm.password
        });
        if (response.status === 200) {
          if (response.data.ok) {
            const username = this.loginForm.username;
            this.$router.push({ name: 'Home', params: { username } });
          } else {
              ElMessageBox.alert('登录失败: ' + response.data.error);
          }
        } else {
          ElMessageBox.alert('登录失败: ' + response.status);
        }
      } catch (error) {
        ElMessageBox.alert('登录失败: ' + error);
      }
    },
    getRouterData() {
      const username = this.$route.params.username;
      if (!!username) {
        this.loginForm.username = username;
      }
    },
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f5f5;
}

.login-title {
  font-size: 36px;
  margin-bottom: 40px;
}

.login-form {
  width: 300px;
}

.error-message {
  color: red;
  margin-top: 10px;
}
</style>
