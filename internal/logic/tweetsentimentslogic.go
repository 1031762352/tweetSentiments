// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TweetSentimentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTweetSentimentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TweetSentimentsLogic {
	return &TweetSentimentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TweetSentimentsLogic) TweetSentiments(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
