## Golang基于Grpc框架搭建

> ### 参考文档
- [proto3](https://developers.google.com/protocol-buffers/docs/proto3)
- [grpc](https://grpc.io/docs/quickstart/)

> ### 生成go的protobuf文件
```
// 根目录下执行
protoc --proto_path=/Users/sky/Code/local/go/src --proto_path=./protos/users  --go_out=plugins=grpc:./protos/users ./protos/users/*.proto
protoc --proto_path=./protos/common  --go_out=plugins=grpc:./protos/common ./protos/common/*.proto
```

> ### 运行
```
    go run server/main.go
    go run client/main.go
```
