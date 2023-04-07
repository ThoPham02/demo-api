package svc

import (
	"demo-api/service/topic/api/internal/config"
	"demo-api/service/topic/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	TopicModel          model.TopicsModel
	AccessLogMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		TopicModel: model.NewTopicsModel(sqlx.NewMysql(c.DataSource)),
	}
}
