// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package influencer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tweetSentiments/internal/logic/influencer"
	"tweetSentiments/internal/svc"
)

// 获取所有订阅大V
func GetInfluencersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := influencer.NewGetInfluencersLogic(r.Context(), svcCtx)
		resp, err := l.GetInfluencers()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
