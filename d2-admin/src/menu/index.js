import { uniqueId } from 'lodash'

/**
 * @description 给菜单数据补充上 path 字段
 * @description https://github.com/d2-projects/d2-admin/issues/209
 * @param {Array} menu 原始的菜单数据
 */
function supplementPath (menu) {
  return menu.map(e => ({
    ...e,
    path: e.path || uniqueId('d2-menu-empty-'),
    ...e.children ? {
      children: supplementPath(e.children)
    } : {}
  }))
}

// 菜单
const menu = [
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '文章',
    icon: 'file-text',
    children: [
      { path: '/post/category', title: '文章分类', icon: 'list' },
      { path: '/tag', title: '标签', icon: 'tags' },
      { path: '/post', title: '文章', icon: 'edit' }
    ]
  },
  {
    title: '外观',
    icon: 'film',
    children: [
      { path: '/nav', title: '菜单', icon: 'location-arrow' },
      {
        title: '主题',
        icon: 'tachometer',
        children: [
          { path: '/theme', title: '主题', icon: 'tachometer' }
        ]
      }
    ]
  },
  {
    title: '用户',
    icon: 'user',
    children: [
      { path: '/user', title: '用户信息', icon: 'user-o' },
      { path: '/comment', title: '评论', icon: 'commenting-o' },
      { path: '/page', title: '页面', icon: 'columns' },
      {
        title: '友链',
        icon: 'link',
        children: [
          { path: '/link/category', title: '友链分类', icon: 'list-ul' },
          { path: '/link', title: '友链', icon: 'link' }
        ]
      }
    ]
  },
  {
    title: '系统',
    icon: 'gears',
    children: [
      { path: '/setting', title: '设置', icon: 'cog' },
      { path: '/attachment', title: '附件', icon: 'save' },
      { path: '/doc', title: 'API 文档', icon: 'book' }
      // { path: '/about', title: '关于', icon: 'paper-plane' }
    ]
  }
]

// 顶部菜单栏
export const menuHeader = supplementPath(menu)

// 侧边菜单栏
export const menuAside = supplementPath(menu)
