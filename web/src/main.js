import Vue from 'vue'
import App from './App.vue'
// 路由配置
import router from './router'
import store from './store' // vuex
import './plugins/element.js' // element ui
// 全局样式
import '@/assets/css/normalize.css';
// 请求
import axios from 'axios';
import { config } from '@fullcalendar/vue'

// 挂载到$http
axios.defaults.baseURL='http://127.0.0.1:8080/';
// 请求头加token
axios.interceptors.request.use(config =>{
  // 拦截后必须要return，继续接下的操作
  if(window.sessionStorage.getItem('token')){
    config.headers.token = window.sessionStorage.getItem('token');
  }

  return config;
},err=>{
  // 失败的操作
  return config;
})
//响应拦截
axios.interceptors.response.use(res => {
  // 必须返回res，否则请求处的res会undefined
  console.log(res);
  return res;
}, err => {
  return res;
})



Vue.prototype.$http = axios

Vue.config.productionTip = false


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')