## 项目介绍

Aries 是基于 Gin + Vue + MySQL + H5 开发的现代化博客系统，博客主要分为后台管理端和展示端。

- 后台管理端采用前后端分离的开发模式，Gin 和 Vue 通过 JSON 进行前后端交互。

- 展示端采用传统开发模式，采用 Go Template 模板引擎来渲染页面，有利于 SEO 优化。

## 功能特性

> **后台管理部分功能和展示端大部分功能还未完成，尚不能正式使用！**

1. 文章

   - markdown 编辑器采用  vditor，提升中文 markdown 使用体验
   - 多标签
   - 文章置顶
   - 文章加密
   - 自定义文章链接
   - 文章排序
   - 草稿和回收站
   - 从文件导入文章

   ![37b31a381a9751ee.png](https://ftp.bmp.ovh/imgs/2020/08/37b31a381a9751ee.png)

   ![9614b942037ca282.png](https://ftp.bmp.ovh/imgs/2020/08/9614b942037ca282.png)

2. 外观

   - 自定义菜单

3. 评论

   - 回收站
   - 评论审核
   - 评论回复邮件通知

4. 图床

   - 支持 sm.ms，imgbb 和 腾讯云 cos 存储
   - 多图片上传
   - 图片管理

   ![b4f25e1d52c50052.png](https://ftp.bmp.ovh/imgs/2020/08/b4f25e1d52c50052.png)

   ![](https://ftp.bmp.ovh/imgs/2020/08/e6d838ea5a9759a8.png)

   ![3142f8dd40ce6d30.png](https://ftp.bmp.ovh/imgs/2020/08/3142f8dd40ce6d30.png)

5. 页面

   - 日志管理
   - 图库管理

   ![b94d1484dfedbe5a.png](https://ftp.bmp.ovh/imgs/2020/08/b94d1484dfedbe5a.png)

   ![d8ac1628e32f8017.png](https://ftp.bmp.ovh/imgs/2020/08/d8ac1628e32f8017.png)

6. 友情链接

   - 添加、修改友链
   - 友链分类

   ![1eae5804d99db93e.png](https://ftp.bmp.ovh/imgs/2020/08/1eae5804d99db93e.png)

7. 用户

   - 修改用户信息
   - 修改密码

   ![ce41bb14aa6d1fca.png](https://ftp.bmp.ovh/imgs/2020/08/ce41bb14aa6d1fca.png)

8. 设置

   - 站点信息和 SEO 信息设置

   ![763266bfba1b89a2.png](https://ftp.bmp.ovh/imgs/2020/08/763266bfba1b89a2.png)

   ![dbiyiq.png](https://s1.ax1x.com/2020/08/30/dbiyiq.png)

## 如何在本地运行 Aries

- 运行环境

  > 请确保已安装好以下环境，配置好 go mod 代理和 npm 国内镜像源
  >
  > ​	如何配置 go mod 代理：https://goproxy.io/zh/
  >
  > ​	如何配置 npm 国内镜像源：https://www.cnblogs.com/luckyhui28/p/12268313.html

  Go 1.13 

  Node.js v12

  MySQL 5.7

- 克隆项目代码到本地

  ```shell
  git clone https://e.coding.net/cangmang/aries/aries.git
  ```

- 新建名称为 **aries** 的数据库，注意字符集为 **utf8m64**，字符编码为 **utf8m64_general_ci**

- 修改项目目录下 `config/develop.yaml` 配置文件中的数据库连接密码
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

- 进入项目中的 `d2-admin` 目录，安装 `npm` 依赖
  ```shell
  npm install
  ```

- 待 `npm` 依赖安装完毕后，启动 `Vue`项目：
  ```shell
  npm run serve
  ```

- 待 `gin` 和 `Vue` 项目启动完毕后，在浏览器中访问 `http://localhost:8080` 即可
