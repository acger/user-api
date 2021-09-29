package logic

import (
	"context"
	"encoding/json"
	"github.com/acger/user-svc/userclient"

	"github.com/acger/user-api/internal/svc"
	"github.com/acger/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EditProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) EditProfileLogic {
	return EditProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditProfileLogic) EditProfile(req types.EditReq) (*types.EditRsp, error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	_, err := l.svcCtx.UserSvc.UserUpdate(l.ctx, &userclient.UserUpdateReq{
		Id:       uint64(uid),
		Name:     req.Name,
		Avatar:   req.Avatar,
		Password: req.Password,
	})

	if err != nil {
		return &types.EditRsp{Code: 1}, nil
	}

	return &types.EditRsp{Code: 0}, nil
}
