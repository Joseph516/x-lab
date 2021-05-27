import Vue from 'vue'
import VueRouter from 'vue-router'


// 组件动态导入
const UploadFile = () => import('@/components/common/UploadFile.vue')



Vue.use(VueRouter)

const routes = [
  {
    path:'/',
    component:UploadFile,
  }

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
