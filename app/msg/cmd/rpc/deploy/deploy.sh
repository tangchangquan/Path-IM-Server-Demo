#!/usr/bin/env bash
tag=`date +%Y%m%d%H%M%S`
rm -rf ./bin
cd ..
CGO_ENABLED=1 CC=x86_64-unknown-linux-gnu-gcc CXX=x86_64-unknown-linux-gnu-g++ GOOS=linux GOARCH=amd64 go build -o deploy/bin
cd deploy
docker build --platform linux/amd64 . -t ccr.ccs.tencentyun.com/zeroim/msg-rpc:${tag}
docker push ccr.ccs.tencentyun.com/zeroim/msg-rpc:${tag}
goctl kube deploy \
--image ccr.ccs.tencentyun.com/zeroim/msg-rpc:${tag} \
--limitCpu 100 \
--limitMem 60 \
--maxReplicas 10 \
--minReplicas 1 \
--name msg-rpc \
--namespace zeroim \
-o ./msg-rpc.yaml \
--port 8080 \
--replicas 1 \
--requestCpu 100 \
--requestMem 60 \
--secret registry \
--serviceAccount find-endpoints \
--home ../../../../../goctl/home