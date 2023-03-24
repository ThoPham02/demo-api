package svc

import (
	"demo-api/account/api/internal/config"
	"demo-api/account/api/internal/middleware"
	"demo-api/account/model"
	"demo-api/account/redisCli"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	AccountsModel       model.AccountsModel
	RedisCache          redis.Redis
	AccessLogMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		AccountsModel:       model.NewAccountsModel(sqlx.NewMysql(c.DataSource)),
		RedisCache:          *redisCli.NewRedis(c.RedisCache.Host, c.RedisCache.Port),
		AccessLogMiddleware: middleware.NewAccessLogMiddleware().Handle,
	}
}
