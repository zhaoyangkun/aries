import { request } from '@/api/service' // 引入封装好的 axios

// 注册
export function authRegister (data) {
  return request({
    url: '/auth/register', // 请求地址
    method: 'post', // 请求类型
    data: data // 请求数据
  })
}

// 登录
export function authLogin (data) {
  return request({
    url: '/auth/login', // 请求地址
    method: 'post', // 请求类型
    data: data // 请求数据
  })
}

// 创建验证码
export function createCaptcha () {
  return request({
    url: '/auth/captcha', // 请求地址
    method: 'get' // 请求类型
  })
}

// 忘记密码
export function forgetPwd (data) {
  return request({
    url: '/auth/pwd/forget',
    method: 'post',
    data: data
  })
}

// 重置密码
export function resetPwd (data) {
  return request({
    url: '/auth/pwd/reset',
    method: 'post',
    data: data
  })
}
