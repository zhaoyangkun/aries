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

// 保存 sm.ms 配置
export function saveSmmsSetting (data) {
  return request({
    url: '/sys_setting/pic_bed/smms',
    method: 'post',
    data: data
  })
}

// 保存 imgbb 配置
export function saveImgbbSetting (data) {
  return request({
    url: '/sys_setting/pic_bed/imgbb',
    method: 'post',
    data: data
  })
}

// 保存腾讯云 COS 配置
export function saveTencentCosSetting (data) {
  return request({
    url: '/sys_setting/pic_bed/tencent_cos',
    method: 'post',
    data: data
  })
}

// 保存评论配置
export function saveCommentSetting (data) {
  return request({
    url: '/sys_setting/comment',
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

// 上传图片
export function uploadImgToPicBed (data) {
  return request({
    url: '/img/attachment/upload',
    method: 'post',
    data: data,
    headers: {
      token: localStorage.getItem('token'),
      'Content-Type': 'multipart/form-data',
      Accept: 'application/json'
    }
  })
}
