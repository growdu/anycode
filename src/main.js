import { createApp } from 'vue'
import './style.css'
import router from './router'; // 引入路由文件
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import axios from 'axios';

import App from './App.vue'
import config from './config'

const app = createApp(App)

app.use(ElementPlus)
app.use(router)
// 设置全局的 API URL 前缀
axios.defaults.baseURL = config.apiUrl;
app.config.globalProperties.$axios = axios;
app.mount('#app')