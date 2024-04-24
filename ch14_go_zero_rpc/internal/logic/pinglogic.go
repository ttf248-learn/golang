package logic

import (
	"context"

	"go-base-learning/ch14_go_zero_rpc/ch14_go_zero_rpc"
	"go-base-learning/ch14_go_zero_rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *ch14_go_zero_rpc.Request) (*ch14_go_zero_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &ch14_go_zero_rpc.Response{}, nil
}
