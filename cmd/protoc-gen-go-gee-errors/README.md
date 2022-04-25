# protoc-gen-go-gee-errors

fork kratos error tool

```text
go install github.com/goapt/gee/protoc-gen-go-gee-errors@latest
```

```shell
protoc --proto_path=. \
  --proto_path=./third_party \
  --go_out=paths=source_relative:./api \
  --go-gee-errors_out=paths=source_relative:./api \
  error_reson.proto
```
