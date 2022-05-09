import { request } from '@api/service'

// 获取所有页面
export function getAllPages () {
  return request({
    url: '/all_pages',
    method: 'get'
  })
}

// 分页获取页面
export function getPagesByPage (params) {
  return request({
    url: '/pages',
    method: 'get',
    params: params
  })
}

// 创建页面
export function createPage (data) {
  return request({
    url: '/pages',
    method: 'post',
    data: data
  })
}

// 更新页面
export function updatePage (data) {
  return request({
    url: '/pages',
    method: 'put',
    data: data
  })
}

// 批量删除页面
export function multiDelPages (ids) {
  return request({
    url: `/pages?ids=${ids}`,
    method: 'delete'
  })
}
