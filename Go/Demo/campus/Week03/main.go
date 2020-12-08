package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	group, groupCtx := errgroup.WithContext(ctx)

	server := http.Server{Addr: "127.0.0.1:8081"}

	// 启动http server
	group.Go(func() error {
		fmt.Printf("启动http监听 \n")
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("http server 启动错误 %v",err)
			cancel()
		}
		return nil
	})

	// 监听linux 信号
	group.Go(func() error {
		fmt.Printf("启动linux信号监听 \n")
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		select {
		case sig := <-sigs:
			fmt.Printf("收到 linux 信号: %s\n", sig)
			cancel()
		case <-groupCtx.Done():
			fmt.Printf("关闭linux信号监听\n")
			return groupCtx.Err()
		}

		return nil
	})


	group.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Printf("http 关闭 \n")
			err := server.Shutdown(ctx)
			if err != nil {
				fmt.Printf(" http关闭错误 %v",err)
				return err
			}
		}
		return nil
	})

	err := group.Wait()

	if err !=  nil {
		if errors.Is(err, context.Canceled) {
			fmt.Printf("context 取消")
		} else {
			fmt.Printf("收到错误 %v\n", err)
		}
	}
}
