package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// 设置路由。使用 beego 默认的路由规则
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is v1.")
}
