#!/usr/bin/env bash
docker build --platform linux/amd64 -t showurl/zerobase:v1.0.0-x86_64 . -f zerobase.dockerfile || exit 1
docker build --platform linux/arm64 -t showurl/zerobase:v1.0.0-arm64 . -f zerobase.dockerfile || exit 1
docker push showurl/zerobase:v1.0.0-x86_64 || exit 1
docker push showurl/zerobase:v1.0.0-arm64 || exit 1
docker manifest create showurl/zerobase:v1.0.0 showurl/zerobase:v1.0.0-x86_64 showurl/zerobase:v1.0.0-arm64 || exit 1
docker manifest push showurl/zerobase:v1.0.0 || exit 1
docker manifest create showurl/zerobase:latest showurl/zerobase:v1.0.0-x86_64 showurl/zerobase:v1.0.0-arm64 || exit 1
docker manifest push showurl/zerobase:latest || exit 1
