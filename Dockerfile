ARG BASEIMAGE=registry.cn-hangzhou.aliyuncs.com/kubespace/distroless-static:latest
FROM $BASEIMAGE

COPY go-app /go-app
CMD ["/go-app"]
