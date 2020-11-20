package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:8081")

	if err != nil {
		panic(err.Error())
	}

	var req float32 = 5

	var resSync *float32

	syncCall := client.Go("MathUtil.CalculateCircleArea", req, &resSync, nil)

	replayDone := <-syncCall.Done

	fmt.Println(replayDone)
	fmt.Println(*resSync)
}