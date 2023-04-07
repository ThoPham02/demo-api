package logic

import (
	"context"

	"demo-api/service/topic/api/internal/svc"
	"demo-api/service/topic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicLogic {
	return &GetTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicLogic) GetTopic(req *types.GetTopicRequest) (resp *types.GetTopicResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
