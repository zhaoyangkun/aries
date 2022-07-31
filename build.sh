#!/bin/bash
# 编译 go 项目
echo "########## start build aries ##########"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o aries ./main.go
# 封装 docker 镜像
echo "########## start build docker image ##########"
docker rmi zhaoyangkun/aries:latest
docker rmi registry.cn-hangzhou.aliyuncs.com/zhaoyangkun/aries:latest
docker build -t zhaoyangkun/aries:latest .
docker tag zhaoyangkun/aries:latest registry.cn-hangzhou.aliyuncs.com/zhaoyangkun/aries:latest
# 推送 docker 镜像
echo "########## start push docker image ##########"
docker push zhaoyangkun/aries:latest
docker push registry.cn-hangzhou.aliyuncs.com/zhaoyangkun/aries:latest
echo "########## all jobs are done ##########"