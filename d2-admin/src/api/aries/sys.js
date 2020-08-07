import { request } from '@/api/service'

// 获取配置条目
export function getSysSettingItem (name) {
  return request({
    url: `/sys_setting/items?name=${name}`,
    method: 'get'
  })
}

// 保存网站配置
export function saveSiteSetting (data) {
  return request({
    url: '/sys_setting/site',
    method: 'post',
    data: data
  })
}

// 保存 SMTP 配置
export function saveSMTPSetting (data) {
  return request({
    url: '/sys_setting/smtp',
    method: 'post',
    data: data
  })
}

// 保存图床配置
export function savePicBedSetting (data) {
  return request({
    url: '/sys_setting/pic_bed',
    method: 'post',
    data: data
  })
}

// 测试发送邮件
export function sendTestEmail (data) {
  return request({
    url: '/sys_setting/email/test',
    method: 'post',
    data: data
  })
}

// 获取后台首页数据
export function getAdminIndexData () {
  return request({
    url: '/sys_setting/index_info',
    method: 'get'
  })
}
