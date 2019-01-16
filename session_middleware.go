package very

import (
	"fmt"
	"github.com/go-redis/redis"
	goredis "github.com/go-redis/redis"
)

var SessionMiddleware = func(rds *redis.Client, session ISession) HandlerFunc {
	return func(c *Context) Response {
		c.Session = session.New()
		// 获取session
		token := c.Request.Header.Get("Access-Token")
		if token == "" {
			return c.Fail(40001, "请先登录")
		}
		// 获取 session 信息
		r := rds.Get(token)
		var err error
		if err = r.Err(); err != nil {
			if err == goredis.Nil {
				return c.Fail(40001, "无效Token")
			}
			return c.BusinessError(err)
		}
		// 登录用户解析
		if err := json.Unmarshal([]byte(r.Val()), c.Session); err != nil {
			return c.BusinessError("Unmarshal error:" + err.Error() + fmt.Sprintf("%+v", c.Session))
		}
		c.Next()
		sess, _ := json.Marshal(c.Session)
		rds.Set(token, string(sess), c.Session.Expire())
		return nil
	}
}
