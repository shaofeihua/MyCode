package main

import (
	"io"
	"log"
	"net/http"
	"time" //用于引入读超时和写超时等
)

var mux map[string]func(http.ResponseWriter, *http.Request) //自定义路由。实现例如 URL://XXX 或者 URL://XXX/123 的跳转

func main() {
	server := http.Server{ //自己实现 server
		Addr:        ":8081",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	//定义路由规则
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = sayHello
	mux["/bye"] = sayBye

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//由于是自定义的路由 mux，如果不增加 if 判断，无法识别并调用自定义路由 mux 中已经注册的两条路由规则 "/hello" 和 "/bye"。所以需要在这里增加判断：定义一个变量 h 去调用 w 和 r（因为 sayHello 和 sayBye 中都调用了 w 和 r），判断路由规则是否已经注册。
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "URL: "+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, this is v3")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye,this is v3")
}
