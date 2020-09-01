import { request } from '@api/service'

// 获取所有图库
export function getAllGalleries () {
  return request({
    url: '/all_galleries',
    method: 'get'
  })
}

// 分页获取图库
export function getGalleriesByPage (params) {
  return request({
    url: '/galleries',
    method: 'get',
    params: params
  })
}

// 创建图库
export function createGallery (data) {
  return request({
    url: '/galleries',
    method: 'post',
    data: data
  })
}

// 更新图库
export function updateGallery (data) {
  return request({
    url: '/galleries',
    method: 'put',
    data: data
  })
}

// 批量删除图库
export function multiDelGalleries (ids) {
  return request({
    url: `/galleries?ids=${ids}`,
    method: 'delete'
  })
}
