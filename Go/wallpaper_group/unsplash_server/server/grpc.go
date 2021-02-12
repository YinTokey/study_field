package server

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/service/grpc"
	"log"
	"time"
	"unsplash_server/api"
	proto "unsplash_server/proto"
)

// CustomFlags : 自定义命令行参数
var CustomFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "dbhost",
		Value: "127.0.0.1",
		Usage: "database address",
	},
	cli.StringFlag{
		Name:  "mqhost",
		Value: "127.0.0.1",
		Usage: "mq(rabbitmq) address",
	},
	cli.StringFlag{
		Name:  "cachehost",
		Value: "127.0.0.1",
		Usage: "cache(redis) address",
	},
	cli.StringFlag{
		Name:  "cephhost",
		Value: "127.0.0.1",
		Usage: "ceph address",
	},
}

type GrpcServer struct {
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (s *GrpcServer) Start() {

	fmt.Println("go micro 服务启动")

	service := grpc.NewService(
		micro.Name("go.micro.service.unsplash"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(CustomFlags...),
	)
	//service.Init(server.Address("127.0.0.1:10086"))
	service.Server().Init(
		server.Address("127.0.0.1:10087"),
	)

	proto.RegisterUnPictureServiceHandler(service.Server(), new(api.UnPicture))

	if err := service.Run(); err != nil {
		log.Println(err)
	}

}