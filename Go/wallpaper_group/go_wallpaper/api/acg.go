package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go_wallpaper/protos/acg_server"
	"google.golang.org/grpc"
	"strconv"
)

func FetchAcgList(c *gin.Context) {
	// 基于consul 服务发现获取ip
	target := AcgServiceEndPoint()

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := timestamppb.NewAcgServiceClient(conn)

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	request := &timestamppb.ListRequest{
		PageSize: int32(pageSize),
		Page:     int32(page),
	}

	result, err := client.List(context.Background(), request)

	if err != nil {
		fmt.Println("grpc 请求错误", err)
	}

	//fmt.Println(result)
	c.JSON(200, result)

}

func AcgServiceEndPoint() string {
	agentAddr := fmt.Sprintf("%v", viper.Get("consul.agentAddr"))

	config := consulapi.DefaultConfig()
	config.Address = agentAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}
	var lastIndex uint64

	services, metainfo, err := client.Health().Service("acg.service", "", true, &consulapi.QueryOptions{
		WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
	})
	if err != nil {
		logrus.Warn("error retrieving instances from Consul: %v", err)
	}
	lastIndex = metainfo.LastIndex

	firstService := services[0]

	endPoint := fmt.Sprintf("%v:%v",firstService.Service.Address,firstService.Service.Port)

	for _, service := range services {
		fmt.Println("service.Service.Address:", service.Service.Address, "service.Service.Port:", service.Service.Port)
		//addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = struct{}{}
	}

	return endPoint

}

func FetchAcgRandom(c *gin.Context) {

}
