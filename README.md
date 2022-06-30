````shell
goctl rpc new rpc
touch user.api
goctl api go -api ./user.api -dir .
goctl model mysql ddl -src user.sql -dir . -c 
goctl rpc protoc ./rpc/user.proto --go_out=./rpc/types --go-grpc_out=./rpc/types  --zrpc_out=./rpc

````



问题1：rpc error: code = DeadlineExceeded desc = context deadline exceede
```shell
了解一下全文信息没有 etcd 字样，但是有 clientv3 表示是 etcd 的客户端和版本，还有 tcp 127.0.0.1:2379: connect: connection refused 连接失败和地址，端口号，这个地址和端口号是 etcd  的，连接失败就是这个意思了。应该是没有启动或防火墙的问题了。
```
