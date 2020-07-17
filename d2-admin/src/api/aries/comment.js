import { request } from '@/api/service'

// 获取所有评论
export function getAllComments () {
  return request({
    url: '/all_comments',
    method: 'get'
  })
}

// 分页获取评论
export function getCommentsByPage (params) {
  return request({
    url: '/comments',
    method: 'get',
    params: params
  })
}

// 发布评论
export function addComment (data) {
  return request({
    url: '/comments',
    method: 'post',
    data: data
  })
}

// 更新评论
export function updateComment (data) {
  return request({
    url: '/comments',
    method: 'put',
    data: data
  })
}

// 删除评论
export function deleteComment (id) {
  return request({
    url: `/comments/${id}`,
    method: 'delete'
  })
}

// 批量删除评论
export function multiDelComments (ids) {
  return request({
    url: `/comments?ids=${ids}`,
    method: 'delete'
  })
}
