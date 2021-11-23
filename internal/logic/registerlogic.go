package logic

import (
	"context"
	"github.com/acger/user-svc/userclient"
	"time"

	"github.com/acger/user-api/internal/svc"
	"github.com/acger/user-api/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (*types.RegisterRsp, error) {

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return &types.RegisterRsp{Code: 2}, nil
	}

	r, err := l.svcCtx.UserSvc.UserAdd(l.ctx, &userclient.UserAddReq{
		Account:  req.Account,
		Password: req.Password,
		Name:     req.Name,
		Avatar:   req.Avatar,
	})

	if err != nil {
		logx.Error("user-add-error: ", err.Error())
		return &types.RegisterRsp{Code: 1}, nil
	}

	if r.Code != 0 {
		logx.Error("user-add-error-code: ", r.Code)
		return &types.RegisterRsp{Code: r.Code, Message: r.Message}, nil
	}

	u, err := l.svcCtx.UserSvc.UserInfo(l.ctx, &userclient.UserInfoReq{Id: r.Uid})

	if err != nil {
		logx.Error("get-user-info-error: ", err.Error())
		return &types.RegisterRsp{Code: 1}, nil
	}

	now := time.Now().Unix()
	auth := l.svcCtx.Config.Auth
	jwt, _ := GetJwtToken(auth.AccessSecret, now, auth.AccessExpire, int64(u.User.Id))

	return &types.RegisterRsp{
		Code: 0,
		User: types.User{
			Id:      u.User.Id,
			Name:    u.User.Name,
			Avatar:  u.User.Avatar,
			Account: u.User.Account,
		},
		Token: jwt,
	}, nil
}
