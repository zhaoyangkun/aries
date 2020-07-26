import layoutHeaderAside from '@/layout/header-aside'

// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

/**
 * 在主框架内显示
 */
const frameIn = [
  {
    path: '/',
    redirect: { name: 'index' },
    component: layoutHeaderAside,
    children: [
      // 首页
      {
        path: 'index',
        name: 'index',
        meta: {
          title: '首页',
          auth: true
        },
        component: _import('system/index')
      },
      // 演示页面
      {
        path: 'category',
        name: 'category',
        meta: {
          title: '分类',
          auth: true
        },
        component: _import('aries/category/category')
      },
      {
        path: 'tag',
        name: 'tag',
        meta: {
          title: '标签',
          auth: true
        },
        component: _import('aries/tag/tag')
      },
      {
        path: 'post',
        name: 'post',
        meta: {
          title: '文章',
          auth: true
        },
        component: _import('aries/post/post')
      },
      {
        path: 'comment',
        name: 'comment',
        meta: {
          title: '评论',
          auth: true
        },
        component: _import('aries/comment/comment')
      },
      {
        path: 'setting',
        name: 'setting',
        meta: {
          title: '设置',
          auth: true
        },
        component: _import('aries/system/setting')
      },
      {
        path: 'doc',
        name: 'doc',
        meta: {
          title: 'API 文档',
          auth: true
        },
        component: _import('aries/system/doc')
      },
      // 系统 前端日志
      {
        path: 'log',
        name: 'log',
        meta: {
          title: '前端日志',
          auth: true
        },
        component: _import('system/log')
      },
      // 刷新页面 必须保留
      {
        path: 'refresh',
        name: 'refresh',
        hidden: true,
        component: _import('system/function/refresh')
      },
      // 页面重定向 必须保留
      {
        path: 'redirect/:route*',
        name: 'redirect',
        hidden: true,
        component: _import('system/function/redirect')
      }
    ]
  }
]

/**
 * 在主框架之外显示
 */
const frameOut = [
  // 登录
  {
    path: '/login',
    name: 'login',
    meta: {
      title: '登录',
      auth: false // 表明无需登录验证
    },
    component: _import('aries/auth/login')
  },
  // 注册
  {
    path: '/register',
    name: 'register',
    meta: {
      title: '初始化配置',
      auth: false // 表明无需登录验证
    },
    // component: import('@/views/aries/auth/register')
    component: _import('aries/auth/register')
  },
  // 忘记密码
  {
    path: '/forgetPwd',
    name: 'forgetPwd',
    meta: {
      title: '忘记密码',
      auth: false // 表明无需登录验证
    },
    component: _import('aries/auth/forgetPwd')
  }
]

/**
 * 错误页面
 */
const errorPage = [
  {
    path: '*',
    name: '404',
    component: _import('system/error/404')
  }
]

// 导出需要显示菜单的
export const frameInRoutes = frameIn

// 重新组织后导出
export default [
  ...frameIn,
  ...frameOut,
  ...errorPage
]
