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

// 顶部菜单栏
export const menuHeader = supplementPath([
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '文章管理',
    icon: 'edit',
    children: [
      { path: '/category', title: '分类', icon: 'list' },
      { path: '/tag', title: '标签', icon: 'tags' },
      { path: '/post', title: '文章', icon: 'file-text' }
      // { path: '/editor', title: 'editor', icon: 'edit' }
    ]
  }
])

// 侧边菜单栏
export const menuAside = supplementPath([
  { path: '/index', title: '首页', icon: 'home' },
  {
    title: '文章管理',
    icon: 'edit',
    children: [
      { path: '/category', title: '分类', icon: 'list' },
      { path: '/tag', title: '标签', icon: 'tags' },
      { path: '/post', title: '文章', icon: 'file-text' }
      // { path: '/editor', title: 'editor', icon: 'edit' }
    ]
  }
])
