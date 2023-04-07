package logic

import (
	"context"
	"demo-api/service/account/api/internal/svc"
	"demo-api/service/account/api/internal/types"
	"demo-api/service/account/api/internal/utils"
	"demo-api/service/account/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountUpdateLogic {
	return &AccountUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountUpdateLogic) AccountUpdate(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("Process update account, input: %v", req)
	userName := l.ctx.Value("userName").(string)
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		l.Logger.Errorf("Failed while hashing password: %v", err)
		return
	}
	err = l.svcCtx.AccountsModel.Update(l.ctx, &model.Accounts{
		Name:         userName,
		HashPassword: hashPassword,
		Email:        req.Email,
	})
	if err != nil {
		l.Logger.Errorf("Failed while updating account, error: %v", err)
		return
	}

	return &types.UpdateResponse{
		Name:  userName,
		Email: req.Email,
	}, nil
}
