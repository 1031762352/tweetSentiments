// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package keyword

import (
	"context"

	"tweetSentiments/internal/svc"
	"tweetSentiments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetKeywordsByCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取指定分类的情感关键词列表
func NewGetKeywordsByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKeywordsByCategoryLogic {
	return &GetKeywordsByCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetKeywordsByCategoryLogic) GetKeywordsByCategory(req *types.GetKeywordsByCategoryRequest) (resp *types.KeywordListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
