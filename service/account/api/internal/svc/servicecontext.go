package svc

import (
	"demo-api/service/account/api/internal/config"
	"demo-api/service/account/api/internal/middleware"
	"demo-api/service/account/model"
	"demo-api/service/account/redisCli"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	AccountsModel       model.AccountsModel
	SessionsModel       model.SessionsModel
	RedisCache          redis.Redis
	AccessLogMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		AccountsModel:       model.NewAccountsModel(sqlx.NewMysql(c.DataSource)),
		SessionsModel:       model.NewSessionsModel(sqlx.NewMysql(c.DataSource)),
		RedisCache:          *redisCli.NewRedis(c.RedisCache.Host, c.RedisCache.Port),
		AccessLogMiddleware: middleware.NewAccessLogMiddleware().Handle,
	}
}
