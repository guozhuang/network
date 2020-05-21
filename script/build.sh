#!/usr/bin/env bash
cd /go/src/SayHello/

#新增dev的环境变量
echo "export GIN_MODE=debug" >> ~/.bashrc
source ~/.bashrc
#热更新机制
go build main.go
./main