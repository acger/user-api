package logic

import (
	"context"
	"encoding/json"
	"github.com/acger/user-api/internal/svc"
	"github.com/acger/user-api/internal/types"
	"github.com/acger/user-api/tool"
	"github.com/acger/user-svc/userclient"
	"github.com/tal-tech/go-zero/core/logx"
)

type ProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProfileLogic {
	return ProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProfileLogic) Profile() (*types.ProfileRsp, error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	u, err := l.svcCtx.UserSvc.UserInfo(l.ctx, &userclient.UserInfoReq{Id: uint64(uid)})

	if err != nil {
		return &types.ProfileRsp{
			Code:    1,
			Message: tool.ErrorCode[1],
		}, nil
	}

	return &types.ProfileRsp{
		Code: 0,
		User: types.User{
			Id:      u.User.Id,
			Name:    u.User.Name,
			Avatar:  u.User.Avatar,
			Account: u.User.Account,
		},
	}, nil
}
