package imuser

import (
	"net/http"

	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/api/internal/logic/imuser"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/api/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/im-user/cmd/api/internal/types"
	"github.com/showurl/Zero-IM-Server/common/xhttp"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserByUsernameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserByUsernameReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			xhttp.ParamErrorResult(r, w, err)
			return
		}

		l := imuser.NewGetUserByUsernameLogic(r.Context(), svcCtx)
		resp, err := l.GetUserByUsername(&req)
		if err != nil {
			xhttp.ParamErrorResult(r, w, err)
			//httpx.Error(w, err)
		} else {
			xhttp.HttpResult(r, w, resp, err)
		}
	}
}
