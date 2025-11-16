// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package sentiment

import (
	"context"

	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTweetSentimentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取大V推文+情感分析结果（支持情感标签筛选）
func NewGetTweetSentimentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTweetSentimentListLogic {
	return &GetTweetSentimentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTweetSentimentListLogic) GetTweetSentimentList(req *types.GetTweetSentimentListRequest) (resp *types.TweetSentimentListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
