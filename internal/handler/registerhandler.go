package handler

import (
	"net/http"

	"github.com/acger/user-api/internal/logic"
	"github.com/acger/user-api/internal/svc"
	"github.com/acger/user-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func registerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
