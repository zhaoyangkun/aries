#!/bin/bash
# 编译 go 项目
echo "########## start build aries ##########"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o aries ./main.go
# 封装 docker 镜像
echo "########## start build docker image ##########"
docker rmi zhaoyangkun/aries:latest
docker build -t zhaoyangkun/aries:latest .
# 推送 docker 镜像
echo "########## start push docker image ##########"
docker push zhaoyangkun/aries:latest
echo "########## all jobs are done ##########"