````shell
goctl rpc new rpc
touch user.api
goctl api go -api ./user.api -dir .
goctl model mysql ddl -src user.sql -dir . -c 
goctl rpc protoc ./rpc/user.proto --go_out=./rpc/types --go-grpc_out=./rpc/types  --zrpc_out=./rpc

````



问题1：rpc error: code = DeadlineExceeded desc = context deadline exceede
```shell
由api 配置文件的consul写法错误有关
```
