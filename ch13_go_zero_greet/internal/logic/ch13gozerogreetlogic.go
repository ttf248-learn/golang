package logic

import (
	"context"

	"go-base-learning/ch13_go_zero_greet/internal/svc"
	"go-base-learning/ch13_go_zero_greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Ch13_go_zero_greetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCh13_go_zero_greetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Ch13_go_zero_greetLogic {
	return &Ch13_go_zero_greetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Ch13_go_zero_greetLogic) Ch13_go_zero_greet(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
