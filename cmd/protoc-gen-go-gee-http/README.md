# protoc-gen-go-very

从 protobuf 文件中生成使用 very 的 http rpc 服务


## 安装

```shell
go install git.verystar.cn/gopkg/protoc-gen-go-very
```

## 生成

```shell
protoc --proto_path=. \
--proto_path=./third_party \
--go-very_out=paths=source_relative:. \
./example/blog/v1/blog.proto
```

## 添加额外tag工具，暂时用不到
```shell
go install github.com/srikrsna/protoc-gen-gotag
```
