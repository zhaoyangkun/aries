import { request } from '@api/service'

// 获取所有日志
export function getAllJournals () {
  return request({
    url: '/all_journals',
    method: 'get'
  })
}

// 分页获取日志
export function getJournalsByPage (params) {
  return request({
    url: '/journals',
    method: 'get',
    params: params
  })
}

// 创建日志
export function createJournal (data) {
  return request({
    url: '/journals',
    method: 'post',
    data: data
  })
}

// 创建日志
export function updateJournal (data) {
  return request({
    url: '/journals',
    method: 'put',
    data: data
  })
}

// 批量删除日志
export function multiDelJournals (ids) {
  return request({
    url: `/journals?ids=${ids}`,
    method: 'delete'
  })
}
