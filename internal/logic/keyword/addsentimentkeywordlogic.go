// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package keyword

import (
	"context"

	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSentimentKeywordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加领域情感关键词（如crypto相关的bullish、bearish）
func NewAddSentimentKeywordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSentimentKeywordLogic {
	return &AddSentimentKeywordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSentimentKeywordLogic) AddSentimentKeyword(req *types.AddKeywordRequest) (resp *types.AddKeywordResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
