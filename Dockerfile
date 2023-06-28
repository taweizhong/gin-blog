FROM golang:latest
LABEL authors="taweizhong"
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/gin-blog
COPY . $GOPATH/src/gin-blog
RUN go mod tidy \
    && go build .
EXPOSE 8000
ENTRYPOINT ["./gin-blog"]


