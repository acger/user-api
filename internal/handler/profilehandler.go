package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"github.com/acger/user-api/internal/logic"
	"github.com/acger/user-api/internal/svc"
)

func profileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewProfileLogic(r.Context(), ctx)
		resp, err := l.Profile()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
