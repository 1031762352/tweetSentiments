// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package influencer

import (
	"context"

	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfluencerByUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过用户名获取大V详情
func NewGetInfluencerByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfluencerByUsernameLogic {
	return &GetInfluencerByUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfluencerByUsernameLogic) GetInfluencerByUsername(req *types.GetInfluencerByUsernameRequest) (resp *types.InfluencerDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
