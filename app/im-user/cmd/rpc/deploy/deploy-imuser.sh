#!/usr/bin/env bash
tag=`date +%Y%m%d%H%M%S`
rm -rf ./bin
cd ..
GOOS=linux GOARCH=amd64 go build -o deploy/bin ./imuser.go
cd deploy
docker build --platform linux/amd64 . -t ccr.ccs.tencentyun.com/zeroim/imuser-rpc:${tag}
docker push ccr.ccs.tencentyun.com/zeroim/imuser-rpc:${tag}
goctl kube deploy \
--image ccr.ccs.tencentyun.com/zeroim/imuser-rpc:20220522153308 \
--limitCpu 100 \
--limitMem 60 \
--maxReplicas 10 \
--minReplicas 1 \
--name imuser-rpc \
--namespace zeroim \
-o ./imuser-rpc.yaml \
--port 8080 \
--replicas 1 \
--requestCpu 100 \
--requestMem 60 \
--secret registry \
--home ../../../../../goctl/home