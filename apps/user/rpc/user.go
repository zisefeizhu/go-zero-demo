package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"go-zero-demo/apps/user/rpc/internal/config"
	"go-zero-demo/apps/user/rpc/internal/server"
	"go-zero-demo/apps/user/rpc/internal/svc"
	"go-zero-demo/apps/user/rpc/types/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
)

//var configFile = flag.String("f", "etc/user.yaml", "the config file")

//
var configFile = flag.String("f", "/Users/zisefeizhu/linkun/goproject/go-zero-demo/apps/user/rpc/etc/user.yaml", "the config file")

/*
   需求改： 增加列出所有用户的接口
*/

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.Timeout = int64(5 * time.Minute)
	ctx := svc.NewServiceContext(c)

	svr := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	err := consul.RegisterService(c.ListenOn, c.Consul)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
