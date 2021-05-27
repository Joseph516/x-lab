/*
 * @Author: yaoyaoyu
 * @Date: 2021-05-27 21:50:55
 */
import Vue from 'vue'
import Element from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import {MessageBox,Message} from 'element-ui'


Vue.use(Element)

// 需要挂载在全局，来使用
Vue.prototype.$message = Message
Vue.prototype.$confirm = MessageBox.confirm

