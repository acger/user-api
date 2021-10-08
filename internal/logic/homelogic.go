package logic

import (
	"context"

	"github.com/acger/user-api/internal/svc"
	"github.com/acger/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeLogic {
	return HomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeLogic) Home() (*types.Rsp, error) {
	return &types.Rsp{Code: 0, Message: "home page."}, nil
}
