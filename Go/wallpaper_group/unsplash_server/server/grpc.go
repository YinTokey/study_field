package server

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
	"unsplash_server/api"
	pb "unsplash_server/proto"
)

const (
	consulAddress = "127.0.0.1:8500"
	localIp       = "127.0.0.1"
	localPort     = 10087
	localServiceName = "service.unsplash"
)

type GrpcServer struct {
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (s *GrpcServer) Start() {

	fmt.Println("grpc 服务启动")

	s.GrpcRegister()

	s.ConsulRegister()

}

// 注册到consul
func (s *GrpcServer) ConsulRegister () {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	//registration.ID = "337"
	registration.Name = localServiceName
	registration.Port = localPort
	//registration.Tags = []string{"testService"}
	registration.Address = localIp

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println(err)
	}
}

// grpc 注册
func (s *GrpcServer) GrpcRegister() {
	addr := fmt.Sprintf("%v:%v",localIp,localPort)
	fmt.Println(addr)
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