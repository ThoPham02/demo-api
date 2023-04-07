package main

import (
	"demo-api/config"
	accountAPI "demo-api/service/account/api"

	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("server_config", "etc/server.yaml", "the config file")

// @BasePath  /api
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	accountService := accountAPI.NewAccountService(server)
	accountService.Start()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
