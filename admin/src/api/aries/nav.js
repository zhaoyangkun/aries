import { request } from '@api/service'

// 获取菜单
export function getNavs () {
  return request({
    url: '/navs',
    method: 'get'
  })
}

// 添加菜单
export function addNav (data) {
  return request({
    url: '/navs',
    method: 'post',
    data: data
  })
}

// 修改菜单
export function updateNav (data) {
  return request({
    url: '/navs',
    method: 'put',
    data: data
  })
}

// 向上移动菜单
export function moveNavUp (navType, orderId) {
  return request({
    url: `/navs/${navType}/up/${orderId}`,
    method: 'patch'
  })
}

// 向下移动菜单
export function moveNavDown (navType, orderId) {
  return request({
    url: `/navs/${navType}/down/${orderId}`,
    method: 'patch'
  })
}

// 删除菜单
export function deleteNav (id) {
  return request({
    url: `/navs/${id}`,
    method: 'delete'
  })
}

// 批量删除菜单
export function multiDelNavs (ids) {
  return request({
    url: `/navs?ids=${ids}`,
    method: 'delete'
  })
}
