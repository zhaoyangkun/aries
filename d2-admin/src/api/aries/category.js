import { request } from '@/api/service'

// 获取所有分类
export function getAllCategories () {
  return request({
    url: '/categories', // 请求地址
    method: 'get' // 请求类型
  })
}
