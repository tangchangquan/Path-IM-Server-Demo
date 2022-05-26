#!/usr/bin/env bash
tag=`date +%Y%m%d%H%M%S`
rm -rf ./bin
cd ..
GOOS=linux GOARCH=amd64 go build -o deploy/bin
cd deploy
docker build --platform linux/amd64 . -t ccr.ccs.tencentyun.com/zeroim/msgtransfer-persistent:${tag}
docker push ccr.ccs.tencentyun.com/zeroim/msgtransfer-persistent:${tag}
#goctl kube deploy \
#--image ccr.ccs.tencentyun.com/zeroim/msgtransfer-persistent:${tag} \
#--limitCpu 100 \
#--limitMem 60 \
#--maxReplicas 10 \
#--minReplicas 1 \
#--name msgtransfer-persistent \
#--namespace zeroim \
#-o ./msgtransfer-persistent.yaml \
#--port 80 \
#--replicas 1 \
#--requestCpu 100 \
#--requestMem 60 \
#--secret registry \
#--serviceAccount find-endpoints \
#--home ../../../../../goctl/home