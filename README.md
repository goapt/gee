<p align="center">
<img src="gee.png"/>
<br/>
<a href="https://github.com/goapt/gee/actions"><img src="https://github.com/goapt/gee/workflows/build/badge.svg" alt="Build Status"></a>
<a href="https://codecov.io/gh/goapt/gee"><img src="https://codecov.io/gh/goapt/gee/branch/master/graph/badge.svg" alt="codecov"></a>
<a href="https://goreportcard.com/report/github.com/goapt/gee"><img src="https://goreportcard.com/badge/github.com/goapt/gee" alt="Go Report Card
"></a>
<a href="https://pkg.go.dev/github.com/goapt/gee"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square" alt="GoDoc"></a>
<a href="https://opensource.org/licenses/mit-license.php" rel="nofollow"><img src="https://badges.frapsoft.com/os/mit/mit.svg?v=103"></a>
</p>

<h3 align="center">Gee is base on gin framework</h3>

## Usage

```go
import "github.com/goapt/gee"

router := gee.Default()

router.GET("/", func(c *gee.Context) error {
    return c.String("hello")
})

```


## Proto API tools

```
go install github.com/goapt/gee/cmd/protoc-gen-go-gee-errors
go install github.com/goapt/gee/cmd/protoc-gen-go-gee-http
```

generate error
```shell
    protoc --proto_path=. --proto_path=../third_party \
    --go_out=paths=source_relative:. \
    --go-gee-errors_out=paths=source_relative:. \
    ./proto/demo/v1/error_reason.proto
```

generate http
```shell
    protoc --proto_path=. --proto_path=../third_party \
    --go_out=paths=source_relative:. \
    --go-gee-http_out=paths=source_relative:. \
    ./proto/demo/v1/demo.proto
```

For example: `/example`
