// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package tweet

import (
	"context"

	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfluencerTweetsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取大V原始推文（不包含情感分析结果）
func NewGetInfluencerTweetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfluencerTweetsLogic {
	return &GetInfluencerTweetsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfluencerTweetsLogic) GetInfluencerTweets(req *types.GetInfluencerTweetsRequest) (resp *types.TweetListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
