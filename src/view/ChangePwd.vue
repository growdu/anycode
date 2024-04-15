<template>
  <div class="change-container">
    <h1 class="change-title">AnyCode</h1>
    <el-form ref="changeForm" :model="changeForm" :rules="changeRules" label-width="80px" label-position="top" class="change-form">
      <el-form-item label="用户名" prop="username" style="align-items: flex-start;">
        <el-input v-model="changeForm.username" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="原密码" prop="password" style="align-items: flex-start;">
        <el-input type="password" v-model="changeForm.password" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="新密码" prop="newpassword" style="align-items: flex-start;">
        <el-input type="password" v-model="changeForm.newpassword" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="确认密码" prop="repassword" style="align-items: flex-start;">
        <el-input type="password" v-model="changeForm.repassword" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="change" style="width: 100%">修改密码</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      changeForm: {
        username: '',
        password: '',
        repassword: ''
      },
      changeRules: {
        username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
        password: [{ required: true, message: '请输入原密码', trigger: 'blur' }],
        newpassword: [{ required: true, message: '请输入新密码', trigger: 'blur' }],
        repassword: [{ required: true, message: '请再次输入新密码', trigger: 'blur' }]
      }
    };
  },
  created() {
    this.getRouterData();
  },
  methods: {
    change() {
      // 在这里获取用户名，并通过路由参数传递给 home 界面
      const username =this.changeForm.username;
      this.$router.push({ name: 'LoginWithParam', params: { username } });
    },
    getRouterData() {
      this.changeForm.username = this.$route.params.username;
      console.log(this.changeForm.username);
    }
  }
};
</script>

<style scoped>
.change-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f5f5;
}

.change-title {
  font-size: 36px;
  margin-bottom: 40px;
}

.change-form {
  width: 300px;
}
</style>
