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

// 获取大V推文+情感分析结果（支持情感标签筛选）
func GetTweetSentimentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTweetSentimentListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sentiment.NewGetTweetSentimentListLogic(r.Context(), svcCtx)
		resp, err := l.GetTweetSentimentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
