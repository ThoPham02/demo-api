package handler

import (
	"net/http"

	"demo-api/account/api/internal/logic"
	"demo-api/account/api/internal/svc"
	"demo-api/account/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AccountLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		
		l := logic.NewAccountLoginLogic(r.Context(), svcCtx)
		resp, err := l.AccountLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
