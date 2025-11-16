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

// 获取指定分类的情感关键词列表
func GetKeywordsByCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetKeywordsByCategoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := keyword.NewGetKeywordsByCategoryLogic(r.Context(), svcCtx)
		resp, err := l.GetKeywordsByCategory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
