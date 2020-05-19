FROM golang:1.13
MAINTAINER guo.zhuang
COPY . /$GOPATH/src/SayHello/
WORKDIR /$GOPATH/src/SayHello/
#设置环境变量，开启go module和设置下载代理
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

#设置业务开发的配置环境
RUN go env -w DEBUG=1
EXPOSE 9000:8080
#CMD ["go","run","main.go"]
CMD ["/bin/bash", "/go/src/SayHello/script/build.sh"]