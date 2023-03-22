package svc

import (
	"demo-api/account/api/internal/config"
	"demo-api/account/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	AccountsModel model.AccountsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AccountsModel: model.NewAccountsModel(sqlx.NewMysql(c.DataSource)),
	}
}
