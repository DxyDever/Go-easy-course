package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(writer, request)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, foo)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func main() {

	//foo这是一个函数名，只不过这个函数是下面的这个形式
	/*
		func(ResponseWriter, *Request)
	*/

	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":8080", nil)
}
