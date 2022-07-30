# 构建 alpine 镜像
FROM alpine

# 声明工作目录
WORKDIR /src/aries

# 复制二进制和静态资源文件
COPY ./aries /src/aries/
COPY ./resources/ /src/aries/resources/

# 声明时区
ENV TZ=Asia/Shanghai

# 解决时区问题
RUN echo "https://mirrors.aliyun.com/alpine/v3.4/main/" > /etc/apk/repositories \
    && apk --no-cache add tzdata zeromq \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo '$TZ' > /etc/timezone

# 向其他容器暴露 8088 端口
EXPOSE 8088

# 命令行运行 gin 项目（需先在根目录下编译好 gin 项目，生成二进制文件）
# 预先编译命令： CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o aries ./main.go
CMD ["./aries"]




