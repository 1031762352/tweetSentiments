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

// 暂停/恢复大V推文拉取（0=暂停，1=恢复）
func UpdateInfluencerActiveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateInfluencerActiveRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := influencer.NewUpdateInfluencerActiveLogic(r.Context(), svcCtx)
		resp, err := l.UpdateInfluencerActive(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
