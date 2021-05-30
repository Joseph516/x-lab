import Vue from 'vue'
import VueRouter from 'vue-router'


// 组件动态导入
const NotFullCalendar = ()=> import('@/components/common/NotFullCalendar.vue')
const UploadFile = () => import('@/components/common/UploadFile.vue')
const Home = ()=> import('@/views/home/Home.vue')
const HomeHeader = ()=> import('@/components/content/HomeHeader.vue')
//  登录
const Login = ()=> import('@/views/login/Login.vue')

const Event = ()=> import('@/components/common/Event.vue')

Vue.use(VueRouter)

const routes = [
  { path:'/', component:UploadFile,},
  { path:'/login', component:Login},
  { path:'/home', component:Home,},
  {
    path:'/homeheader',
    component:HomeHeader
  },
  {
    path:'/notfullcalendar',
    component:NotFullCalendar,
  },
  {path:'/event',component:Event},

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
