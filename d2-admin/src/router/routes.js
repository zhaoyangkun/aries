import layoutHeaderAside from '@/layout/header-aside'

import Index from '@/views/aries/system/index'

import PostCategory from '@/views/aries/category/postCategory'

import Tag from '@/views/aries/tag/tag'

import Post from '@/views/aries/post/post'

import Comment from '@/views/aries/comment/comment'

import Page from '@/views/aries/page/page'

import Journal from '@/views/aries/page/journal'

import Gallery from '@/views/aries/page/gallery'

import User from '@/views/aries/user/user'

import Setting from '@/views/aries/system/setting'

import Links from '@/views/aries/link/links'

import LinkCategory from '@/views/aries/category/linkCategory'

import Navs from '@/views/aries/nav/navs'

import Attachment from '@/views/aries/system/attachment'

import Theme from '@/views/aries/theme/theme'

import Doc from '@/views/aries/system/doc'

import Log from '@/views/system/log'

import Refresh from '@/views/system/function/refresh'

import Redirect from '@/views/system/function/redirect'

import Login from '@/views/aries/auth/login'

import Register from '@/views/aries/auth/register'

import ForgetPwd from '@/views/aries/auth/forgetPwd'

import ResetPwd from '@/views/aries/auth/resetPwd'

import Error404 from '@/views/system/error/404'

// 由于懒加载页面太多的话会造成 webpack 热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
// const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

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
        component: Index
      },
      {
        path: 'post/category',
        name: 'postCategory',
        meta: {
          title: '文章分类',
          auth: true
        },
        component: PostCategory
      },
      {
        path: 'tag',
        name: 'tag',
        meta: {
          title: '标签',
          auth: true
        },
        component: Tag
      },
      {
        path: 'post',
        name: 'post',
        meta: {
          title: '文章',
          auth: true
        },
        component: Post
      },
      {
        path: 'comment',
        name: 'comment',
        meta: {
          title: '评论',
          auth: true
        },
        component: Comment
      },
      {
        path: 'page',
        name: 'page',
        meta: {
          title: '页面',
          auth: true
        },
        component: Page
      },
      {
        path: 'page/journal',
        name: 'journal',
        meta: {
          title: '日志',
          auth: true
        },
        component: Journal
      },
      {
        path: 'page/gallery',
        name: 'gallery',
        meta: {
          title: '图库',
          auth: true
        },
        component: Gallery
      },
      {
        path: 'user',
        name: 'user',
        meta: {
          title: '用户信息',
          auth: true
        },
        component: User
      },
      {
        path: 'setting',
        name: 'setting',
        meta: {
          title: '设置',
          auth: true
        },
        component: Setting
      },
      {
        path: 'link',
        name: 'links',
        meta: {
          title: '友链',
          auth: true
        },
        component: Links
      },
      {
        path: 'link/category',
        name: 'linkCategory',
        meta: {
          title: '友链分类',
          auth: true
        },
        component: LinkCategory
      },
      {
        path: 'nav',
        name: 'navs',
        meta: {
          title: '菜单',
          auth: true
        },
        component: Navs
      },
      {
        path: 'attachment',
        name: 'attachment',
        meta: {
          title: '附件',
          auth: true
        },
        component: Attachment
      },
      {
        path: 'theme',
        name: 'theme',
        meta: {
          title: '主题',
          auth: true
        },
        component: Theme
      },
      {
        path: 'doc',
        name: 'doc',
        meta: {
          title: 'API 文档',
          auth: true
        },
        component: Doc
      },
      // 系统 前端日志
      {
        path: 'log',
        name: 'log',
        meta: {
          title: '前端日志',
          auth: true
        },
        component: Log
      },
      // 刷新页面 必须保留
      {
        path: 'refresh',
        name: 'refresh',
        hidden: true,
        component: Refresh
      },
      // 页面重定向 必须保留
      {
        path: 'redirect/:route*',
        name: 'redirect',
        hidden: true,
        component: Redirect
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
    component: Login
  },
  // 注册
  {
    path: '/register',
    name: 'register',
    meta: {
      title: '初始化配置',
      auth: false // 表明无需登录验证
    },
    component: Register
  },
  // 忘记密码
  {
    path: '/forgetPwd',
    name: 'forgetPwd',
    meta: {
      title: '忘记密码',
      auth: false // 表明无需登录验证
    },
    component: ForgetPwd
  },
  // 重置密码
  {
    path: '/resetPwd',
    name: 'resetPwd',
    meta: {
      title: '忘记密码',
      auth: false // 表明无需登录验证
    },
    component: ResetPwd
  }
]

/**
 * 错误页面
 */
const errorPage = [
  {
    path: '*',
    name: '404',
    component: Error404
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
