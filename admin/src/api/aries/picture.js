import { request } from '@api/service'

// 分页获取图片
export function getImagesByPage (params) {
  return request({
    url: '/images',
    method: 'get',
    params: params
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
