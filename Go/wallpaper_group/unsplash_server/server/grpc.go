
package server

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
	"unsplash_server/internal/routes/api"
	pb "unsplash_server/proto"
)

const (
	consulAddress = "127.0.0.1:8500"
	localIp       = "127.0.0.1"
	localPort     = 10088
	localServiceName = "service.unsplash"
)

type GrpcServer struct {
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (s *GrpcServer) Start() {

	//s.ConsulRegister()

	s.GrpcRegister()

}

// 注册到consul
func (s *GrpcServer) ConsulRegister () {
	fmt.Println("consul 开始服务注册")
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "unsplash_1"
	registration.Name = localServiceName
	registration.Port = localPort
	registration.Tags = []string{"unsplash"}
	registration.Address = localIp

	// 增加consul健康检查回调函数
	registration.Check = &consulapi.AgentServiceCheck{ // 健康检查
		//HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",  // 健康检查间隔
		DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务，注销时间，相当于过期时间
		GRPC:     fmt.Sprintf("%v:%v/%v", localIp, localPort, localServiceName),// grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
	}

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println("consul 注册失败 ", err)
	}
	fmt.Println("consul 注册完成 ")

}

// grpc 注册
func (s *GrpcServer) GrpcRegister() {
	addr := fmt.Sprintf("%v:%v",localIp,localPort)
	fmt.Println("gprc 服务注册 ", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterUnPictureServiceServer(server, &api.UnsplashServer{})
	if err := server.Serve(lis); err != nil {
		fmt.Println("failed to serve: %v", err)
	}

}

// 注册到 go-micro
func (s *GrpcServer) RegToGoMicro() {

	//reg := consul.NewRegistry()
	//
	//service := grpc.NewService(
	//	micro.Name("go.micro.service.unsplash"),
	//	micro.RegisterTTL(time.Second*10),
	//	micro.RegisterInterval(time.Second*5),
	//	micro.Flags(CustomFlags...),
	//	micro.Registry(reg),
	//)
	////service.Init(server.Address("127.0.0.1:10086"))
	//service.Server().Init(
	//	server.Address("127.0.0.1:10087"),
	//)
	//
	//proto.RegisterUnPictureServiceHandler(service.Server(), new(api.UnPicture))
	//
	//if err := service.Run(); err != nil {
	//	log.Println(err)
	//}
}

