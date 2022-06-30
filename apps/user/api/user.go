package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"go-zero-demo/apps/user/api/internal/config"
	"go-zero-demo/apps/user/api/internal/handler"
	"go-zero-demo/apps/user/api/internal/svc"
	"time"
)

//var configFile = flag.String("f", "etc/user.yaml", "the config file")

var configFile = flag.String("f", "/Users/zisefeizhu/linkun/goproject/go-zero-demo/apps/user/api/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.Timeout = int64(5 * time.Minute)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
