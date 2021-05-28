import Vue from 'vue'
import App from './App.vue'
// 路由配置
import router from './router'
import store from './store' // vuex
import './plugins/element.js' // element ui

// 全局样式
import '@/assets/css/normalize.css';

Vue.config.productionTip = false


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')