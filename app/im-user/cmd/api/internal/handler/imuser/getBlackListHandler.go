package imuser

import (
	"net/http"

	"github.com/showurl/Path-IM-Server/app/im-user/cmd/api/internal/logic/imuser"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/api/internal/svc"
	"github.com/showurl/Path-IM-Server/app/im-user/cmd/api/internal/types"
	"github.com/showurl/Path-IM-Server/common/xhttp"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBlackListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBlackListReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			xhttp.ParamErrorResult(r, w, err)
			return
		}

		l := imuser.NewGetBlackListLogic(r.Context(), svcCtx)
		resp, err := l.GetBlackList(&req)
		if err != nil {
			xhttp.ParamErrorResult(r, w, err)
			//httpx.Error(w, err)
		} else {
			xhttp.HttpResult(r, w, resp, err)
		}
	}
}
