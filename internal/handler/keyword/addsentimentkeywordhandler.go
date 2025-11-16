// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package keyword

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tweetSentiments/internal/logic/keyword"
	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"
)

// 添加领域情感关键词（如crypto相关的bullish、bearish）
func AddSentimentKeywordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddKeywordRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := keyword.NewAddSentimentKeywordLogic(r.Context(), svcCtx)
		resp, err := l.AddSentimentKeyword(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
