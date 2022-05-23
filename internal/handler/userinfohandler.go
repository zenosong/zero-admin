package handler

import (
	"net/http"

	"zero-admin/internal/logic"
	"zero-admin/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func userInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
