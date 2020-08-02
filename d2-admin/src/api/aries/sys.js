import { request } from '@/api/service'

// 获取配置条目
export function getSysSettingItem (name) {
  return request({
    url: `/sys_setting_items?name=${name}`,
    method: 'get'
  })
}

// 保存网站配置
export function saveSiteSetting (data) {
  return request({
    url: '/site_setting',
    method: 'post',
    data: data
  })
}

// 保存 SMTP 配置
export function saveSMTPSetting (data) {
  return request({
    url: '/smtp_setting',
    method: 'post',
    data: data
  })
}

// 测试发送邮件
export function sendTestEmail (data) {
  return request({
    url: '/test_send_email',
    method: 'post',
    data: data
  })
}
