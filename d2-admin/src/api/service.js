import axios from 'axios'
import Adapter from 'axios-mock-adapter'
import { get } from 'lodash'
import { errorCreate, errorLog } from './tools'
import router from '@/router'
import Vue from 'vue'

/**
 * @description 创建请求实例
 */
function createService () {
  // 创建一个 axios 实例
  const service = axios.create()
  // 请求拦截
  service.interceptors.request.use(
    config => config,
    error => {
      // 发送失败
      console.log(error)
      return Promise.reject(error)
    }
  )
  // 响应拦截
  service.interceptors.response.use(
    response => {
      // dataAxios 是 axios 返回数据中的 data
      const dataAxios = response.data
      // 这个状态码是和后端约定的
      const code = dataAxios.code
      // 根据 code 进行判断
      if (code === undefined) {
        // 如果没有 code 代表这不是项目后端开发的接口 比如可能是 D2Admin 请求最新版本
        return dataAxios
      } else {
        // 有 code 代表这是一个后端接口 可以进行进一步的判断
        switch (code) {
          // code === 100 代表请求成功
          case 100:
            return dataAxios
          // code === 101 代表重定向
          case 101:
            router.replace({
              path: `${dataAxios.data.url}`
            }).then(r => {
            })
            break
          // code === 102 表示禁止访问
          case 102:
            // 显示错误信息
            Vue.prototype.$message.error(dataAxios.msg, { duration: 1200 })
            // 跳转到登录页面
            router.replace({
              path: '/login'
            }).then(r => {
            })
            break
          // code === 103 或 104 代表有错误发生
          default:
            errorCreate(`${dataAxios.msg}`)
            break
        }
      }
    },
    error => {
      const status = get(error, 'response.status')
      switch (status) {
        case 400:
          error.message = '请求错误'
          break
        case 401:
          error.message = '未授权，请登录'
          break
        case 403:
          error.message = '拒绝访问'
          break
        case 404:
          error.message = `请求地址出错: ${error.response.config.url}`
          break
        case 408:
          error.message = '请求超时'
          break
        case 500:
          error.message = '服务器端错误'
          break
        case 501:
          error.message = '服务未实现'
          break
        case 502:
          error.message = '网关错误'
          break
        case 503:
          error.message = '服务不可用'
          break
        case 504:
          error.message = '网关超时'
          break
        case 505:
          error.message = 'HTTP版本不受支持'
          break
        default:
          break
      }
      errorLog(error)
      return Promise.reject(error)
    }
  )
  return service
}

/**
 * @description 创建请求方法
 * @param {Object} service axios 实例
 */
function createRequestFunction (service) {
  return function (config) {
    const token = localStorage.getItem('token')
    const configDefault = {
      headers: {
        // token
        token: token,
        // 请求数据类型
        'Content-Type': get(config, 'headers.Content-Type', 'application/json'),
        // 返回数据类型
        Accept: 'application/json'
      },
      timeout: 50000,
      baseURL: process.env.VUE_APP_API,
      data: {}
    }
    return service(Object.assign(configDefault, config))
  }
}

// 用于真实网络请求的实例和请求方法
export const service = createService()
export const request = createRequestFunction(service)

// 用于模拟网络请求的实例和请求方法
export const serviceForMock = createService()
export const requestForMock = createRequestFunction(serviceForMock)

// 网络请求数据模拟工具
export const mock = new Adapter(serviceForMock)
