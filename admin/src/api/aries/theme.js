import { request } from '@/api/service'

// 获取所有主题
export function getAllThemes () {
  return request({
    url: '/all_themes',
    method: 'get'
  })
}

// 根据名称获取主题
export function getThemeByName (name) {
  return request({
    url: `/themes/${name}`,
    method: 'get'
  })
}

// 启用主题
export function enableTheme (data) {
  return request({
    url: '/themes',
    method: 'post',
    data: data
  })
}
