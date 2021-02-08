# 选择一个镜像，可以通过命令 docker images 查看本地下载的镜像
FROM golang:latest

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 运行终端命令，创建文件夹
RUN mkdir -p /usr/src/app

# 移动到工作目录：
WORKDIR /usr/src/app

ADD go.mod .
ADD go.sum .

RUN export GO111MODULE=on
RUN export GOPROXY=https://goproxy.cn
RUN go mod download
RUN go get -u github.com/go-sql-driver/mysql

# 将工程代码复制到容器中
COPY . /usr/src/app

RUN go build -o app .


# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令
CMD ["./app"]
