FROM debian:buster-slim
WORKDIR /app
ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive


RUN sed -i 's#http://deb.debian.org#http://mirrors.aliyun.com#g' /etc/apt/sources.list && \
    ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && dpkg-reconfigure --frontend noninteractive tzdata && \
    set -x
RUN apt-get update && bash -c "sleep 10 && apt-get update"
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates apt-transport-https procps net-tools vim curl tzdata gawk && \
    rm -rf /var/lib/apt/lists/*

# docker build --platform linux/amd64 -t showurl/zerobase:v1.0.0-x86_64 . -f zerobase.dockerfile
# docker build --platform linux/arm64 -t showurl/zerobase:v1.0.0-arm64 . -f zerobase.dockerfile
# docker push showurl/zerobase:v1.0.0-x86_64
# docker push showurl/zerobase:v1.0.0-arm64
# docker manifest create showurl/zerobase:v1.0.0 showurl/zerobase:v1.0.0-x86_64 showurl/zerobase:v1.0.0-arm64
# docker manifest push showurl/zerobase:v1.0.0
# docker manifest create showurl/zerobase:latest showurl/zerobase:v1.0.0-x86_64 showurl/zerobase:v1.0.0-arm64
# docker manifest push showurl/zerobase:latest
