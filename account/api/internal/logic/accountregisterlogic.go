package logic

import (
	"context"

	"demo-api/account/api/internal/svc"
	"demo-api/account/api/internal/types"
	"demo-api/account/api/internal/utils"
	"demo-api/account/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountRegisterLogic {
	return &AccountRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountRegisterLogic) AccountRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("Process register account, input: %v", req)

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		l.Logger.Errorf("Failed while hashing password: %v", err)
		return nil, err
	}

	_, err = l.svcCtx.AccountsModel.Insert(l.ctx, &model.Accounts{
		Name:         req.Name,
		HashPassword: hashPassword,
		Email:        req.Email,
	})
	if err != nil {
		l.Logger.Errorf("Failed while creating account: %v", err)
		return nil, err
	}

	return &types.RegisterResponse{
		Name:  req.Name,
		Email: req.Email,
	}, nil
}
