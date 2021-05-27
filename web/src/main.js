import Vue from 'vue'
import App from './App.vue'
// 路由配置
import router from './router'
import store from './store' // vuex
import './plugins/element.js' // element ui


// axios ajax
import axios from 'axios'
// baseUrl
axios.defaults.baseURL = 'http://localhost:8080/';
// 请求拦截
axios.interceptors.request.use(config => {
  return config;
}, err => {
  // 失败的操作
})
//  响应拦截
axios.interceptors.response.use(res => {
  console.log('res :>> ', res);
  return res
}, err => {
  return res
})



Vue.config.productionTip = false
// 挂载
Vue.prototype.$http = axios

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')