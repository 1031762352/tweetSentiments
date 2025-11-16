// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package influencer

import (
	"context"

	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInfluencerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加订阅大V（通过Twitter用户名）
func NewAddInfluencerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInfluencerLogic {
	return &AddInfluencerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddInfluencerLogic) AddInfluencer(req *types.AddInfluencerRequest) (resp *types.InfluencerDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
