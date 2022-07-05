# protoc-gen-go-gee-http

从 protobuf 文件中生成使用 gee 的 http rpc 服务


## 安装

```shell
go install github.com/goapt/gee/cmd/protoc-gen-go-gee-http@latest
```

## 生成

```shell
protoc --proto_path=. \
--proto_path=./third_party \
--go-gee-http_out=paths=source_relative:. \
./example/blog/v1/blog.proto
```

## 添加额外tag工具，暂时用不到
```shell
go install github.com/srikrsna/protoc-gen-gotag
```
