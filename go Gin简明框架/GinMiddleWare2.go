package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexFunc(c *gin.Context) {
	fmt.Println("index")
	v, ok := c.Get("name")
	if ok {
		fmt.Println(v)
	}
	c.String(http.StatusOK, "index")
}

//中间件1
func Middleware1() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("m1 before")
		c.Set("name", "周杰伦")
		c.Next()
		fmt.Println("m1 after")
	}
}

//中间件2
func Middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("m1 before")
		c.Set("name", "周杰伦")
		c.Next()
		fmt.Println("m1 after")
	}
}

func main() {

	r := gin.Default()

	//全局注册
	//r.Use(Middleware1(), Middleware2())

	//路由组注册
	//r.Group("/shop", Middleware1(), Middleware2())

	//注册中间件
	r.GET("/index", Middleware1(), Middleware2(), indexFunc)

	r.Run()

}