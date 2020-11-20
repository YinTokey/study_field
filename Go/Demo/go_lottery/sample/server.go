package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct {
}

func (mu *MathUtil) CalculateCircleArea(req float32, resp *float32) error {
	fmt.Println("远程调用")
	*resp = req * req //圆形的面积 s = π * r * r
	return nil
}


func main() {
	mathUtil := new(MathUtil) //初始化指针数据类型

	//2、调用net/rpc包的功能将服务对象进行注册
	err := rpc.Register(mathUtil)
	if err != nil {
		panic(err.Error())
	}

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(listen, nil)
}