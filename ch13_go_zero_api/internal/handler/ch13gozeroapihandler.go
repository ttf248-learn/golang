package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-base-learning/ch13_go_zero_api/internal/logic"
	"go-base-learning/ch13_go_zero_api/internal/svc"
	"go-base-learning/ch13_go_zero_api/internal/types"
)

func Ch13_go_zero_apiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCh13_go_zero_apiLogic(r.Context(), svcCtx)
		resp, err := l.Ch13_go_zero_api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
