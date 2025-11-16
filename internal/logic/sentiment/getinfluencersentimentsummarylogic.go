// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package sentiment

import (
	"context"

	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfluencerSentimentSummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取大V推文情感分布和平均得分
func NewGetInfluencerSentimentSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfluencerSentimentSummaryLogic {
	return &GetInfluencerSentimentSummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfluencerSentimentSummaryLogic) GetInfluencerSentimentSummary(req *types.GetInfluencerSentimentSummaryRequest) (resp *types.InfluencerSentimentSummaryResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
