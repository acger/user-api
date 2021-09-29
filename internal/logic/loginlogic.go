package logic

import (
	"context"
	"github.com/acger/user-api/tool"
	"github.com/acger/user-svc/userclient"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/acger/user-api/internal/svc"
	"github.com/acger/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func GetJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginRsp, error) {
	us := l.svcCtx.UserSvc

	u, err := us.UserInfo(l.ctx, &userclient.UserInfoReq{
		Account: req.Account,
	})

	if err != nil {
		return &types.LoginRsp{
			Code:    11001,
			Message: tool.ErrorCode[11001],
		}, nil
	}

	if u.Code != 0 {
		return &types.LoginRsp{
			Code:    u.Code,
			Message: u.Message,
		}, nil
	}

	pwdErr := bcrypt.CompareHashAndPassword([]byte(u.User.Password), []byte(req.Password))
	if pwdErr != nil {
		return &types.LoginRsp{
			Code:    11002,
			Message: tool.ErrorCode[11002],
		}, nil
	}

	now := time.Now().Unix()
	auth := l.svcCtx.Config.Auth
	jwt, _ := GetJwtToken(auth.AccessSecret, now, auth.AccessExpire, int64(u.User.Id))

	return &types.LoginRsp{
		Code: 0,
		User: types.User{
			Id:      u.User.Id,
			Account: u.User.Account,
			Name:    u.User.Name,
			Avatar:  u.User.Avatar,
		},
		Token: jwt,
	}, nil
}
