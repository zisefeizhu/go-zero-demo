package handler

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/apps/user/api/internal/logic"
	"go-zero-demo/apps/user/api/internal/svc"
	"go-zero-demo/apps/user/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		fmt.Println(r.Context().Err())
		// bug: 这里出现 context deadline exceeded
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		fmt.Println(r.Context().Err())
		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
