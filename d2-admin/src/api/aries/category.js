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

// 添加文章分类
export function addArticleCategory (data) {
  return request({
    url: '/categories/article',
    method: 'post',
    data: data
  })
}

// 修改文章分类
export function updateArticleCategory (data) {
  return request({
    url: '/categories/article',
    method: 'put',
    data: data
  })
}

// 添加友链分类
export function addLinkCategory (data) {
  return request({
    url: '/categories/link',
    method: 'post',
    data: data
  })
}

// 修改友链分类
export function updateLinkCategory (data) {
  return request({
    url: '/categories/link',
    method: 'put',
    data: data
  })
}

// 添加图库分类
export function addGalleryCategory (data) {
  return request({
    url: '/categories/gallery',
    method: 'post',
    data: data
  })
}

// 修改图库分类
export function updateGalleryCategory (data) {
  return request({
    url: '/categories/gallery',
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
