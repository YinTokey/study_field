package main

import (
	"go_wallpaper/message"

	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

type OrderServiceImpl struct {
}

//具体的方法实现
//GetOrderInfo(context.Context, *OrderRequest) (*OrderInfo, error)
func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {

	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}

	var response *message.OrderInfo
	current := time.Now().Unix()
	if (request.TimeStamp > current) {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			fmt.Println(result)
			return &result, nil
		} else {
			return nil, errors.New("server error")
		}
	}

	return response, nil
}


func main() {

	server := grpc.NewServer()

	message.RegisterOrderServiceServer(server,new(OrderServiceImpl))

	lis, err := net.Listen("tcp","localhost:8091")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}