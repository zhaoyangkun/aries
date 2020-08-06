import { request } from '@api/service'

// 获取所有友链
export function getAllLinks () {
  return request({
    url: '/all_links',
    method: 'get'
  })
}

// 分页获取友链
export function getLinksByPage (params) {
  return request({
    url: '/links',
    method: 'get',
    params: params
  })
}

// 添加友链
export function createLink (data) {
  return request({
    url: '/links',
    method: 'post',
    data: data
  })
}

// 更新友链
export function updateLink (data) {
  return request({
    url: '/links',
    method: 'put',
    data: data
  })
}

// 删除友链
export function deleteLink (id) {
  return request({
    url: `/links/${id}`,
    method: 'delete'
  })
}

// 批量删除友链
export function multiDelLinks (ids) {
  return request({
    url: `/links?ids=${ids}`,
    method: 'delete'
  })
}
