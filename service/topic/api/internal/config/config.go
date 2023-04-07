package config

import (
	"time"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DataSource string
	Auth       struct {
		AccessSecret  string
		AccessExpire  time.Duration
		RefreshExpire time.Duration
	}
}
