package handler

import (
	"net/http"

	"demo-api/service/topic/api/internal/logic"
	"demo-api/service/topic/api/internal/svc"
	"demo-api/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTopicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTopicRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetTopicLogic(r.Context(), svcCtx)
		resp, err := l.GetTopic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
