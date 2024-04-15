<template>
  <div class="home">
    <h1 class="title">Welcome, {{ username }}</h1>
    <el-table :data="tableData" style="width: 100%">
      <el-table-column prop="name" label="名称"></el-table-column>
      <el-table-column prop="connectionString" label="vscode地址"></el-table-column>
      <el-table-column prop="sshPort" label="ssh端口"></el-table-column>
      <el-table-column prop="rootPassword" label="root密码"></el-table-column>
      <el-table-column prop="coderPassword" label="coder密码"></el-table-column>
      <el-table-column label="Status">
        <template #default="{ row }">
          <el-badge :value="row.status" :class="statusClass(row.status)" :status="statusIcon(row.status)">
          </el-badge>
        </template>
      </el-table-column>
    </el-table>
    <div class="button-group">
      <el-button type="primary" @click="start">启动</el-button>
      <el-button type="danger" @click="stop">停止</el-button>
      <el-button type="warning" @click="restart">重启</el-button>
      <el-button type="success" @click="startProgramming">开始编程</el-button>
      <el-button type="success" @click="update">升级anycode</el-button>
      <el-button type="warning" @click="change">修改密码</el-button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      tableData: [
        { name: 'Server 1', connectionString: 'example1.com', sshPort: 22, rootPassword: '123', coderPassword: '123', status: 'stop' }
      ]
    };
  },
  created() {
    this.getRouterData();
  },
  methods: {
    getRouterData() {
      this.username = this.$route.params.username;
    },
    start() {
      console.log('Starting server...');
    },
    stop() {
      console.log('Stopping server...');
    },
    restart() {
      console.log('Restarting server...');
    },
    startProgramming() {
      console.log('Starting programming...');
    },
    update() {
      console.log('Updating server...');
    },
    change() {
      console.log('Change user password...');
      const username = this.username;
      this.$router.push({ name: 'Change', params: { username } });
    },
    statusClass(status) {
      return {
        'is-success': status === 'running',
        'is-danger': status === 'stop',
        'is-warning': status === 'unknown'
      };
    },
    statusIcon(status) {
      return {
        'success': status === 'running',
        'error': status === 'stop',
        'warning': status === 'unknown'
      };
    }
  }
};
</script>

<style scoped>
.home {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: #f8f8f8;
}

.title {
  font-size: 24px;
  margin-bottom: 20px;
}

.button-group {
  margin-top: 20px;
}
</style>
