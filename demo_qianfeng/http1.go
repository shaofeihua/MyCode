/*
	静态 web 服务器开发：
		1、定义 url 对应的处理器/处理器函数
		2、启动 web 服务
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.Handle("/static/", http.FileServer(http.Dir(".")))
	addr := ":9999"
	http.Handle("/static/", NewStaticHandler("."))
	http.ListenAndServe(addr, nil)
}

// 定义一个处理器
type StaticHandler struct {
	dir string // 定义静态服务器的目录
}

func NewStaticHandler(dir string) *StaticHandler {
	return &StaticHandler{
		dir: dir,
	}
}

// 处理器要实现的方法
func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取用户请求的 URL 使用参数 r
	fmt.Println(r.URL)
}
