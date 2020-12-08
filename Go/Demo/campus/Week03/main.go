package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// 注册通知
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <- sigs
		fmt.Println(sig)
		done <- true

	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

	s1 := http.Server{Addr: "8080"}
	s2 := http.Server{Addr: "8081"}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	group, _ := errgroup.WithContext(ctx)

	group.Go(func() error {
		if err := s1.ListenAndServe(); err != nil {
			cancel()
			return err
		}
		return nil
	})

	group.Go(func() error {
		if err := s2.ListenAndServe(); err != nil {
			cancel()
			return err
		}
		return nil
	})

}
