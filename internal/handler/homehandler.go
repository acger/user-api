package handler

import (
	"net/http"

	"github.com/acger/user-api/internal/logic"
	"github.com/acger/user-api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func homeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewHomeLogic(r.Context(), ctx)
		resp, err := l.Home()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
