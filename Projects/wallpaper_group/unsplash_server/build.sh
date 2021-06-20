# 编译成linux下可运行的程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o unsplash-app
