package main

import (
	"demo-api/service/topic/api/internal/config"
	"demo-api/service/topic/api/internal/handler"
	"demo-api/service/topic/api/internal/svc"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

type TopicService struct {
	C      config.Config
	Server *rest.Server
	Ctx    *svc.ServiceContext
}

func NewTopicService(server *rest.Server) *TopicService {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	handler.RegisterHandlers(server, ctx)

	return &TopicService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

func (as *TopicService) Start() error {
	return nil
}
