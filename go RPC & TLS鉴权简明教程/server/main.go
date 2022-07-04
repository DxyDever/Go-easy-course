package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {

	rpc.Register(new(Cal))

	rpc.HandleHTTP()

	log.Printf("Serving RPC server on port %d", 1234)

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)

	}
}

/*
使用 rpc.Register，发布 Cal 中满足 RPC 注册条件的方法（Cal.Square）
使用 rpc.HandleHTTP 注册用于处理 RPC 消息的 HTTP Handler
使用 http.ListenAndServe 监听 1234 端口，等待 RPC 请求。

*/
