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



]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})


router.beforeEach( (to,from,next) => {
  // 添加过滤
  if(to.path === '/login') return next();
  // 从sessionstorage中取到保存token值
  const tokenStr = window.sessionStorage.getItem('token')
  //没有token则跳转到登录页
  // todo 这里应该做axios请求去校验
  if(!tokenStr) return next('/login')
  next()
})



export default router
