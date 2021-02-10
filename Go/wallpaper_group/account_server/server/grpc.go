package server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (s *GrpcServer) Start() {

	fmt.Println("grpc 服务启动中。。")

	//configs.Init()

	server := grpc.NewServer()

	//api.RegisterUnPictureServiceServer(server, new(UnPictureServiceImpl))

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("grpc 服务启动开始")

	server.Serve(lis)

}