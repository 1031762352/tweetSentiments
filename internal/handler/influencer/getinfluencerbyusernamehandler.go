// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package influencer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tweetSentiments/internal/logic/influencer"
	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"
)

// 通过用户名获取大V详情
func GetInfluencerByUsernameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInfluencerByUsernameRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := influencer.NewGetInfluencerByUsernameLogic(r.Context(), svcCtx)
		resp, err := l.GetInfluencerByUsername(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
