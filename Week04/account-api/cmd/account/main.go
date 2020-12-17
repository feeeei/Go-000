package main

import (
	api_v1 "account/api/v1"
	"account/internal/service"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		return serveRPC(ctx)
	})
	group.Go(func() error {
		return serveSignal(ctx)
	})

	if err := group.Wait(); err != nil {
		shutdown()
	}
}

func serveRPC(ctx context.Context) error {
	grpcServer := grpc.NewServer()
	api_v1.RegisterHeartbeatServer(grpcServer, service.NewHeartbeat())
	lis, err := net.Listen("tcp", "9000")
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		grpcServer.Stop()
	}()
	return grpcServer.Serve(lis)
}

func serveSignal(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-c:
		return fmt.Errorf("Receive shutdown signal")
	case <-ctx.Done():
		return nil
	}
}

func shutdown() {
	finish := make(chan struct{})
	go func() {
		log.Println("Starting shutdown...")
		// close_conn
		// flush_log
		// clean_cache
		// ......
		time.Sleep(time.Millisecond * 200) // 模拟后台处理收尾任务200ms
		log.Println("Finished shutdown")
		finish <- struct{}{}
	}()
	select {
	case <-finish:
		return
	case <-time.After(time.Second * 1):
		return
	}
}
