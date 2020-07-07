## 如何运行 Aries

- 运行环境

  > 请确保已安装好以下环境，配置好 go mod 代理和 npm 国内镜像源
  >
  > ​	如何配置 go mod 代理：https://goproxy.io/zh/
  >
  > ​	如何配置 npm 国内镜像源：https://www.cnblogs.com/luckyhui28/p/12268313.html

  - Go 1.13 

  - Node.js v12

  - MySQL 5.7

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
