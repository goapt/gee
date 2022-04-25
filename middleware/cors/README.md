# CORS


# Use

```go

type Cors very.HandlerFunc

func NewCors() Cors {
	return cors.New(cors.Config{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"*"},
        AllowHeaders: []string{
        "Origin",
        "Content-Length",
        "Content-Type",
        "Access-Token",
        "Access-Control-Allow-Origin",
        "x-requested-with",
        },
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
	})
}
```