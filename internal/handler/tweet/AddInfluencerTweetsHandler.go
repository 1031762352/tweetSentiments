package tweet

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tweetSentiments/internal/logic/tweet"
	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"
)

// 添加大V原始推文（不包含情感分析结果）
func AddInfluencerTweetsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInfluencerTweetsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tweet.NewGetInfluencerTweetsLogic(r.Context(), svcCtx)
		resp, err := l.GetInfluencerTweets(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
