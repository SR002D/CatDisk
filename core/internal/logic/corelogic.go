package logic

import (
	"context"
	"fmt"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
	ans := fmt.Sprintf("Hello %s!", req.Name)
	resp = &types.Response{
		ans,
	}

	return
}