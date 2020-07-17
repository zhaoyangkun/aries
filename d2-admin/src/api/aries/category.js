import { request } from '@/api/service'

// 获取所有分类
export function getAllCategories (type) {
  return request({
    url: `/all_categories?category_type=${type}`, // 请求地址
    method: 'get' // 请求类型
  })
}

// 分页获取分类
export function getCategoriesByPage (params) {
  return request({
    url: '/categories',
    method: 'get',
    params: params
  })
}

// 获取所有父级分类
export function getAllParentCategories (type) {
  return request({
    url: `/parent_categories?category_type=${type}`,
    method: 'get'
  })
}

// 添加分类
export function addCategory (data) {
  return request({
    url: '/categories',
    method: 'post',
    data: data
  })
}

// 修改分类
export function updateCategory (data) {
  return request({
    url: '/categories',
    method: 'put',
    data: data
  })
}

// 删除分类
export function deleteCategory (id) {
  return request({
    url: `/categories/${id}`,
    method: 'delete'
  })
}

// 批量删除分类
export function multiDelCategories (ids) {
  return request({
    url: `/categories?ids=${ids}`,
    method: 'delete'
  })
}
