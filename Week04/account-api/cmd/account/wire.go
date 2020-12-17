package main

import (
	"account/internal/service"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitRPCServe() *service.Heartbeat {
	wire.Build(grpc.NewServer(), service.NewHeartbeat())
	return &service.Heartbeat{}
}
