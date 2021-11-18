/*
	多路复用器
	Go 语言中有默认的多路复用器，不用自己写
*/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := NewMux()
	mux.Handle("/a.xxx", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("a.xxxxx")
		fmt.Fprintln(w, "a.xxxx")
	}))

	addr := ":8888"
	http.ListenAndServe(addr, mux)
}

// 定义一个多路复用器
type Mux struct {
	routers map[string]http.Handler
}

func NewMux() *Mux {
	return &Mux{
		// 路由初始化
		routers: map[string]http.Handler{},
	}
}

// 所有的 url 处理最终都会经过 ServeHTTP()
func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	// 在路由里面找请求的 url
	if handler, ok := mux.routers[r.URL.Path]; ok {
		// 如果处理器 handler 里面包含 url 的处理关系
		handler.ServeHTTP(w, r)
	} else {
		// 如果处理器 handler 里面不包含 url 的处理关系
		w.WriteHeader(http.StatusNotFound)
	}
}

// 定义 url 和处理器的关系：注册路由
func (mux *Mux) Handle(url string, h http.Handler) {
	mux.routers[url] = h
}
