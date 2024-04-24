package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-base-learning/ch13_go_zero_greet/internal/logic"
	"go-base-learning/ch13_go_zero_greet/internal/svc"
	"go-base-learning/ch13_go_zero_greet/internal/types"
)

func Ch13_go_zero_greetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCh13_go_zero_greetLogic(r.Context(), svcCtx)
		resp, err := l.Ch13_go_zero_greet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
