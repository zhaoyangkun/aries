import { request } from '@/api/service'

// 获取所有用户
export function getAllUsers () {
  return request({
    url: '/all_users',
    method: 'get'
  })
}

// 修改用户信息
export function updateUser (data) {
  return request({
    url: '/users',
    method: 'put',
    data: data
  })
}

// 修改用户密码
export function updateUserPwd (data) {
  return request({
    url: '/users/pwd',
    method: 'put',
    data: data
  })
}
