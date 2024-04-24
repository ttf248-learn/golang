package logic

import (
	"context"

	"go-base-learning/ch13_go_zero_api/internal/svc"
	"go-base-learning/ch13_go_zero_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Ch13_go_zero_apiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCh13_go_zero_apiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Ch13_go_zero_apiLogic {
	return &Ch13_go_zero_apiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Ch13_go_zero_apiLogic) Ch13_go_zero_api(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = "Hello " + req.Name

	return
}
