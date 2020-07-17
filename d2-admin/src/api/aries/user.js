import { request } from '@/api/service'

export function getAllUsers () {
  return request({
    url: '/all_users',
    method: 'get'
  })
}
