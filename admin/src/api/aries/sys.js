import { request } from '@/api/service'

// 获取博客全局变量
export function getBlogVars () {
  return request({
    url: '/sys_setting/blog_vars',
    method: 'get'
  })
}

// 获取配置条目
export function getSysSettingItem (name) {
  return request({
    url: `/sys_setting/items?name=${name}`,
    method: 'get'
  })
}

// 获取后台首页数据
export function getAdminIndexData () {
  return request({
    url: '/sys_setting/index_info',
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

// 测试发送邮件
export function sendTestEmail (data) {
  return request({
    url: '/sys_setting/email/test',
    method: 'post',
    data: data
  })
}

// 保存去不图床配置
export function saveQubuSetting (data) {
  return request({
    url: '/sys_setting/pic_bed/qubu',
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

// 保存本地评论配置
export function saveLocalCommentSetting (data) {
  return request({
    url: '/sys_setting/comment/local',
    method: 'post',
    data: data
  })
}

// 保存 Twikoo 组件设置
export function saveTwikooSetting (data) {
  return request({
    url: '/sys_setting/comment/twikoo',
    method: 'post',
    data: data
  })
}

// 保存参数配置
export function saveParamSetting (data) {
  return request({
    url: '/sys_setting/param',
    method: 'post',
    data: data
  })
}

// 保存社交信息
export function saveSocialInfo (data) {
  return request({
    url: '/sys_setting/social_info',
    method: 'post',
    data: data
  })
}
