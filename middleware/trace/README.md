# opentelemetry traceing middleware


# Use

```go

type Trace gee.Hanler

func NewTrace(conf *config.Config) Trace {
	tr := trace.New(trace.WithAppName(conf.App.AppName))
	return tr
}
```
