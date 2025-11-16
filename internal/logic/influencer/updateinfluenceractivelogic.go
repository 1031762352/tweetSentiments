// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package influencer

import (
	"context"

	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfluencerActiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 暂停/恢复大V推文拉取（0=暂停，1=恢复）
func NewUpdateInfluencerActiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfluencerActiveLogic {
	return &UpdateInfluencerActiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInfluencerActiveLogic) UpdateInfluencerActive(req *types.UpdateInfluencerActiveRequest) (resp *types.InfluencerDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
