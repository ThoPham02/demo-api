package logic

import (
	"context"
	"demo-api/service/account/api/internal/svc"
	"demo-api/service/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountGetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountGetInfoLogic {
	return &AccountGetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountGetInfoLogic) AccountGetInfo(req *types.InfoRequest) (resp *types.InfoResponse, err error) {
	l.Logger.Infof("Process Get info of account %v", req)

	userName := l.ctx.Value("userName").(string)
	account, err := l.svcCtx.AccountsModel.FindOne(l.ctx, userName)
	if err != nil {
		l.Logger.Errorf("Failed while find account by name: %s", userName)
		return nil, err
	}

	resp = &types.InfoResponse{
		Name:  account.Name,
		Email: account.Email,
	}

	return
}
