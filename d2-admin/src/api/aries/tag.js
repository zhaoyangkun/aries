import { request } from '@/api/service'

// 获取所有标签
export function getAllTags () {
  return request({
    url: '/all_tags',
    method: 'get'
  })
}

// 分页获取标签
export function getTagsByPage (params) {
  return request({
    url: '/tags',
    method: 'get',
    params: params
  })
}

// 根据 ID 获取标签
export function getTagById (id) {
  return request({
    url: `/tags/${id}`,
    method: 'get'
  })
}

// 添加标签
export function addTag (data) {
  return request({
    url: '/tags',
    method: 'post',
    data: data
  })
}

// 更新标签
export function updateTag (data) {
  return request({
    url: '/tags',
    method: 'put',
    data: data
  })
}

// 删除标签
export function deleteTag (id) {
  return request({
    url: `/tags/${id}`,
    method: 'delete'
  })
}

// 批量删除标签
export function multiDelTags (ids) {
  return request({
    url: `/tags?ids=${ids}`,
    method: 'delete'
  })
}
