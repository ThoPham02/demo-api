package logic

import (
	"context"
	"errors"
	"time"

	"demo-api/account/api/internal/svc"
	"demo-api/account/api/internal/types"
	"demo-api/account/api/internal/utils"
	"demo-api/account/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLoginLogic {
	return &AccountLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountLoginLogic) AccountLogin(req *types.LoginRequest) (*types.LoginResponse, error) {
	l.Logger.Infof("Process Login Account, input: %v", &req)
	account, err := l.svcCtx.AccountsModel.FindOne(l.ctx, req.Name)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("this account does not exist")
		}
		l.Logger.Errorf("Failed while get account by name, %s", err.Error())
		return nil, err
	}

	err = utils.VerifyPassword(req.Password, account.HashPassword)
	if err != nil {
		l.Logger.Errorf("Failed while verify password, %s", err.Error())
		return nil, errors.New("wrong password")
	}

	secretKey := l.svcCtx.Config.Auth.AccessSecret
	iat := time.Now().Unix()
	seconds := l.svcCtx.Config.Auth.AccessExpire
	accesstoken, err := l.getJwtToken(secretKey, iat, seconds, account.Name)
	if err != nil {
		l.Logger.Errorf("Failed while getting token: %v", err)
		return nil, err
	}

	return &types.LoginResponse{
		AccessToken: accesstoken,
	}, nil
}

// Gen JWT token from secret key, time and user name
func (l *AccountLoginLogic) getJwtToken(secretKey string, iat, seconds int64, userName string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userName"] = userName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
