# 构建 alpine 镜像
FROM alpine

# 声明工作目录
WORKDIR /src/aries

# 将当前目录下所有内容复制到工作目录
COPY . /src/aries

# 向其他容器暴露 8088 端口
EXPOSE 8088

# 命令行运行 gin 项目（需先在根目录下编译好 gin 项目，生成二进制文件）
CMD ["./aries"]




