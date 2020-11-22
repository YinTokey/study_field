package main

import (
	"context"
	"fmt"
	"go_wallpaper/message"
	"google.golang.org/grpc"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:8091",grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	orderRequest := &message.OrderRequest{OrderId: "201908300001", TimeStamp: time.Now().Unix()}
	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)

	if orderInfo != nil {
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
		fmt.Println(orderInfo.GetOrderStatus())
	} else {
		fmt.Println(err.Error())
	}

}