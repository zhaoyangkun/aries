## 项目介绍

> **主题管理**和**评论组件**功能还在开发中！

Aries 是基于 Gin + GORM + MySQL + Vue + H5 开发的现代化博客系统，博客主要分为**后台管理端**和**展示端**。

- **后台管理端**采用**前后端分离**的开发模式，Gin 和 Vue 通过 JSON 进行前后端交互。
- **展示端**采用传统开发模式，采用 Go Template 模板引擎来渲染页面，有利于 SEO 优化。
- 数据库层没有设置物理外键，所有外键操作都在业务层处理。

## 功能特性

1. 文章
   - markdown 编辑器采用  vditor，提升中文 markdown 使用体验
   - 文章标签、分类管理
   - 文章置顶
   - 文章加密
   - 自定义文章链接
   - 文章排序
   - 草稿箱和回收站
   - 从文件导入文章

2. 外观

   - 自定义菜单
   - 导入主题
   - 主题切换和设置
   - 主题编辑
3. 评论

   - 回收站
   - 评论审核
   - 评论回复邮件通知
4. 图床

   - 支持 sm.ms，imgbb 和 腾讯云 cos 存储
   - 多图片上传
   - 图片管理

5. 页面

   - 日志管理
   - 自定义页面
   - 图库管理

6. 友情链接

   - 添加、修改友链
   - 友链分类

7. 用户

   - 修改用户信息
   - 修改密码

8. 设置

   - 网站设置
   - 邮件设置
   - 图床设置
   - 评论设置
   - 参数设置


## 部分截图

- 后台管理端

  ![初始化配置](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910102026.png)

  ![登录](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910101720.png)

  ![写文章](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910102450.png)

  ![设置](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910102605.png)

  ![文章列表](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910111630.png)

  ![图片上传](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910104827.png)

  ![附件列表](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910111523.png)

  ![添加图库](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910111730.png)

- 博客展示端

  ![首页](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910105929.png)

  ![文章详情](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910110249.png)

  ![分类列表](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910110729.png)

  ![标签列表](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910110801.png)

  ![归档](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910110325.png)

  ![友链](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910110634.png)

  ![相册](https://e4dwm6.coding-pages.com/article/2020/9/10/20200910111228.png)

## 如何在本地运行 Aries

- 运行环境：

  > 请确保已安装好以下环境，配置好 go mod 代理和 npm 国内镜像源
  >
  > ​	如何配置 go mod 代理：https://goproxy.io/zh/
  >
  > ​	如何配置 npm 国内镜像源：https://www.cnblogs.com/luckyhui28/p/12268313.html

  Go 1.13 

  Node.js v12

  MySQL 5.7

- 克隆项目代码到本地：

  ```shell
  git clone https://e.coding.net/cangmang/aries/aries.git
  ```

- 新建名称为 **aries** 的数据库，注意字符集为 **utf8m64**，字符编码为 **utf8m64_general_ci**。

- 修改项目目录下 `config/develop.yaml` 配置文件中的数据库连接密码：
  ```yaml
  # 服务器配置
  server:
    mode: "debug"                               # 运行模式
    port: "8088"                                # 端口
    token_expire_time: 3600                     # JWT token 过期时间（单位：秒）
    allowed_refers: ["localhost", "127.0.0.1"]  # 允许的 referer
  # 数据库配置
  db:
    host: "127.0.0.1"       # 主机地址
    user_name: "root"       # 用户名
    password: "19960331"    # 密码
    database: "aries"       # 数据库名
    port: "3306"            # 端口
    time_zone: "Local"      # 时区
    max_idle_conn: 10       # 最大空闲连接数
    max_open_conn: 20       # 最大打开连接数
  ```

- 进入项目根目录，安装 `gin` 项目相关依赖：
  ```shell 
  go mod download
  go get -u github.com/swaggo/swag/cmd/swag 
  ```

- 进入项目根目录，启动 `gin` 项目：
  ```shell
  go run main.go
  ```

- 进入项目中的 `d2-admin` 目录，安装 `npm` 依赖：
  ```shell
  npm install
  ```

- 待 `npm` 依赖安装完毕后，启动 `Vue`项目：
  ```shell
  npm run serve
  ```

- 待 `gin` 和 `Vue` 项目启动完毕后，在浏览器中访问 `http://localhost:8080` 即可进入后台管理。

- 在浏览器中访问 `http://127.0.0.1:8088` 可进入博客展示端。
