// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package influencer

import (
	"context"

	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfluencersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取所有订阅大V
func NewGetInfluencersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfluencersLogic {
	return &GetInfluencersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfluencersLogic) GetInfluencers() (resp *types.InfluencerListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
