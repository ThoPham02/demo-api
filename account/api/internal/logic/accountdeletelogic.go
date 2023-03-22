package logic

import (
	"context"

	"demo-api/account/api/internal/svc"
	"demo-api/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountDeleteLogic {
	return &AccountDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountDeleteLogic) AccountDelete(req *types.DeleteRequest) error {
	// todo: add your logic here and delete this line
	userName := l.ctx.Value("userName").(string)
	l.Logger.Infof("Process delete account, userName: %v", userName)

	err := l.svcCtx.AccountsModel.Delete(l.ctx, userName)
	if err != nil {
		l.Logger.Errorf("Failed while deleting account, error: %v", err)
		return err
	}

	return nil
}
