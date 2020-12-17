// service 层
//    做 DTO -> DO 的转换
//    做 biz 层的编排逻辑
package service

import (
	account_v1 "account/api/v1"
	"context"
)

type Heartbeat struct {
}

func NewHeartbeat() *Heartbeat {
	return &Heartbeat{}
}

func (h *Heartbeat) Heartbeat(ctx context.Context, request *account_v1.Heart) (*account_v1.Heart, error) {
	return &account_v1.Heart{Type: account_v1.HeartbeatType_Pong, Ts: request.Ts}, nil
}
