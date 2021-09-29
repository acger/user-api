package logic

import (
	"context"
	"fmt"
	"github.com/acger/user-api/internal/svc"
	"github.com/acger/user-api/internal/types"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/tal-tech/go-zero/core/logx"
)

type QiniuUpTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQiniuUpTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) QiniuUpTokenLogic {
	return QiniuUpTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QiniuUpTokenLogic) QiniuUpToken(req types.QiniuUpReq) (*types.QiniuUpRsp, error) {
	config := l.svcCtx.Config.Qiniu
	var scope string

	if req.Name != "" {
		scope = fmt.Sprintf("%s:%s", config.Bucket, req.Name)
	} else {
		scope = config.Bucket
	}

	putPolicy := storage.PutPolicy{
		Scope: scope,
	}

	mac := qbox.NewMac(config.AK, config.SK)
	upToken := putPolicy.UploadToken(mac)

	return &types.QiniuUpRsp{Code: 0, Token: upToken}, nil
}
