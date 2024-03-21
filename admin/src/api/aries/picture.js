import { request } from '@api/service'

// const qubuUploadURL = 'https://7bu.top/api/v1/upload'
// const smmsUploadURL = 'https://sm.ms/api/v2/upload'
// const imgbbUploadURL = 'https://api.imgbb.com/1/upload'

// 分页获取图片
export function getImagesByPage (params) {
  return request({
    url: '/images',
    method: 'get',
    params: params
  })
}

// 保存图片信息
export function saveImageInfo (data) {
  return request({
    url: '/images',
    method: 'post',
    data: data
  })
}

// 上传图片到附件
export function uploadImgToAttachment (data, config = {}) {
  return request({
    url: '/images/attachment/upload',
    method: 'post',
    data: data,
    headers: {
      token: localStorage.getItem('token'),
      'Content-Type': 'multipart/form-data',
      Accept: 'application/json'
    },
    ...config
  })
}

// 批量删除图片
export function multiDelImages (ids) {
  return request({
    url: `/images?ids=${ids}`,
    method: 'delete'
  })
}

// 上传图片到 7bu 图床
export function uploadImgTo7bu (data, token, config = {}) {
  return request({
    baseURL: process.env.VUE_APP_7_BU_API,
    url: '/upload',
    method: 'post',
    data: data,
    headers: {
      Authorization: `Bearer ${token}`,
      'Content-Type': 'multipart/form-data',
      Accept: 'application/json'
    },
    ...config
  })
}

// 上传图片到 picui 图床
export function uploadImgToPicUI (data, token, config = {}) {
  return request({
    baseURL: process.env.VUE_APP_PIC_UI_API,
    url: '/upload',
    method: 'post',
    data: data,
    headers: {
      Authorization: `Bearer ${token}`,
      'Content-Type': 'multipart/form-data',
      Accept: 'application/json'
    },
    ...config
  })
}

// 上传图片到 sm.ms 图床
export function uploadImgToSmms (data, token, config = {}) {
  return request({
    baseURL: process.env.VUE_APP_SM_MS_API,
    url: '/upload',
    method: 'post',
    data: data,
    headers: {
      Authorization: token,
      'Content-Type': 'multipart/form-data',
      Accept: 'application/json'
    },
    ...config
  })
}

// 上传图片到 imgbb 图床
export function uploadImgToImgbb (data, token, config = {}) {
  return request({
    baseURL: process.env.VUE_APP_IMGBB_API,
    url: `/upload?key=${token}`,
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    ...config
  })
}
