package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func Logging() Middleware {

	//create a new middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {

			start := time.Now()

			defer func() {
				log.Println("-----------------", request.URL.Path, time.Since(start))
			}()
			log.Println("我是爷爷")
			f(writer, request)
		}
	}
}

func Method(m string) Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {

			if request.Method != m {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			}

			f(writer, request)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}
