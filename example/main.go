package main

import (
	"log"

	"github.com/goapt/gee"
	gerrors "github.com/goapt/gee/errors"
	"github.com/goapt/gee/example/proto/demo/v1"
)

func main() {
	router := gee.New()

	gee.ErrorHandler = func(c *gee.Context, err error) error {
		err2 := gerrors.FromError(err)
		c.Status(int(err2.GetCode()))
		c.Abort()
		return c.JSON(gee.H{
			"code": err2.GetReason(),
			"msg":  err2.GetMessage(),
		})
	}

	router.Use(func(c *gee.Context) error {
		c.Response.Before(func(w *gee.Response) {
			if w.Body() != nil {
				c.Writer.Header().Set("x-body-sign", "before")
			}
		})

		c.Next()
		return nil
	})

	router.Use(func(c *gee.Context) error {
		if c.Query("a") == "2" {
			// return errors.New("a is middleware 2")
		}
		c.Next()

		return nil
	})

	router.GET("/test", func(c *gee.Context) error {
		if c.Query("a") == "2" {
			return demo.ErrorAccessForbidden("No access")
		}

		return c.JSON(gee.H{
			"message": "Hello World!",
		})
	})

	err := router.Run(":9999")
	if err != nil {
		log.Fatal(err)
	}
}
