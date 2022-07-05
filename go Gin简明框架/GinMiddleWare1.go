package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// 默认的路由会自动创建一个Logger（）跟Recovery()两个中间件的router
	//r := gin.Default()

	//创建一个不带默认中间件的router
	r := gin.New()

	//全局注册中间件
	//r.Use(MiddlewareFunc1)

	r.GET("/test1", MiddlewareFunc1, func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "调用中间件1")
	})

	r.GET("/test2", MiddlewareFunc2(true), func(c *gin.Context) {
		fmt.Fprint(c.Writer, "调用中间件2")
	})

	r.Run(":8080")
}

// 下面2个中间件作用：
// 统计处理的延时
// 打印处理的状态

func MiddlewareFunc1(c *gin.Context) {
	t := time.Now()
	fmt.Println("before middleware")

	//设置request变量到context的key里面去，获取通过Get等函数取出
	c.Set("request", "client_request")

	//发送request之前

	c.Next() //  去执行下一个处理函数

	//发送request之后
	//这个c.Write是RespnseWriter,，我们可以获取状态等信息
	status := c.Writer.Status()
	fmt.Println("after middleware,", status)

	t2 := time.Since(t)

	fmt.Println("time:", t2)
}

//中间件2，通过闭包的方式返回一个处理函数，但是可以传入外部变量
func MiddlewareFunc2(debug bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if debug {

		} else {

		}

		t := time.Now()
		fmt.Println("before middleware")

		//设置request变量到context的Key中，通过GET等函数可以取得
		c.Set("request", "client_request")
		//发送request之前
		c.Next()
		//发送request之后
		status := c.Writer.Status()
		fmt.Println("after middleware,", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
