// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/acger/user-api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: registerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: loginHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/profile",
				Handler: profileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/profile/edit",
				Handler: editProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/qiniu/up/token",
				Handler: qiniuUpTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
