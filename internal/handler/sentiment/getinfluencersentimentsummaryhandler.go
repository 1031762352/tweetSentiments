// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package sentiment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tweetSentiments/internal/logic/sentiment"
	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"
)

// 获取大V推文情感分布和平均得分
func GetInfluencerSentimentSummaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInfluencerSentimentSummaryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sentiment.NewGetInfluencerSentimentSummaryLogic(r.Context(), svcCtx)
		resp, err := l.GetInfluencerSentimentSummary(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
