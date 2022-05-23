#!/bin/bash
goctl api plugin -plugin \
goctl-swagger="swagger -filename imuser.json -host 42.194.149.177:10000" -api imuser.api -dir .