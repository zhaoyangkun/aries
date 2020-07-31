import { request } from '@/api/service'

// 获取配置条目
export function getSysSettingItem (name) {
  return request({
    url: `/sys_setting_items?name=${name}`,
    method: 'get'
  })
}

// 保存 SMTP 配置信息
export function saveSMTP (data) {
  return request({
    url: '/smtp',
    method: 'post',
    data: data
  })
}

// 发送测试邮件
export function sendTestEmail (data) {
  return request({
    url: '/test_email',
    method: 'post',
    data: data
  })
}
