#!/bin/bash
# 编译 go 项目
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o aries ./main.go
# 封装 docker 镜像
docker build -t zhaoyangkun/aries .
# 推送 docker 镜像
docker push