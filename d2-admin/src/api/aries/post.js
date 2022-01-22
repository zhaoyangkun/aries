import { request } from '@/api/service'

// 分页获取文章
export function getPostsByPage (params) {
  return request({
    url: '/articles', // 请求地址
    method: 'get', // 请求类型
    params: params
  })
}

// 根据 ID 获取文章
export function getPostById (id) {
  return request({
    url: `/articles/${id}`, // 请求地址
    method: 'get' // 请求类型
  })
}

// 添加文章
export function addPost (data) {
  return request({
    url: '/articles', // 请求地址
    method: 'post', // 请求类型
    data: data
  })
}

// 回收或恢复文章
export function recycleOrRecoverPost (id) {
  return request({
    url: `/articles/recycle/${id}`,
    method: 'patch'
  })
}

// 从文件导入文章
export function importPostFromFiles (data) {
  return request({
    url: '/articles/files',
    method: 'post',
    data: data,
    headers: {
      token: localStorage.getItem('token'),
      'Content-Type': 'multipart/form-data',
      Accept: 'application/json'
    }
  })
}

// 修改文章
export function updatePost (data) {
  return request({
    url: '/articles', // 请求地址
    method: 'put', // 请求类型
    data: data
  })
}

// 删除文章
export function deletePost (id) {
  return request({
    url: `/articles/${id}`,
    method: 'delete'
  })
}

// 批量删除文章
export function multiDelPosts (ids) {
  return request({
    url: `/articles?ids=${ids}`,
    method: 'delete'
  })
}

// 向上移动文章
export function movePostUp (data) {
  return request({
    url: '/articles/up',
    method: 'patch',
    data: data
  })
}

// 向下移动文章
export function movePostDown (data) {
  return request({
    url: '/articles/down',
    method: 'patch',
    data: data
  })
}
