/*
 * @Author: yaoyaoyu
 * @Date: 2021-05-27 21:52:03
 */
import axios from 'axios';

export function request(config) {

  const instance = axios.create({
    baseURL: 'localhost:8080',
    timeout: 5000
  })

  // 1.请求拦截，比如config中的一些信息不符合服务器的要求
  // 2.比如每次发送网络请求时，都希望在界面中显示一个请求的图标
  // 3.某些网络请求。比如登陆（token）必须携带一些特殊的信息
  instance.interceptors.request.use(config => {
    // 拦截后必须要return config，不然啥也没有了
    return config
  }, err => {

  })


  //响应拦截
  instance.interceptors.response.use(res => {
    // 必须返回res，否则请求处的res会undefined
    return res
  }, err => {

  })

  return instance(config)
};