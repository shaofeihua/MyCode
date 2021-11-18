//自定义 mux 和 handler
package main

import (
	"io"
	"log"
	"net/http"
	"os" //程序优化：这里引入 os 包目的是为了实现静态文件的存储（go 自带的功能）
)

// 自己写路由规则
func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{}) //定义一个 handle 注册到 mux 中
	mux.HandleFunc("/Hello", sayHello)

	wd, err := os.Getwd() //使用 Getwd 方法获取当前路径
	if err != nil {
		log.Fatal(err)
	}

	// 使用 http.StripPrefix 去掉前缀
	// 使用 http.FileServer 定义静态文件服务器的路径
	// 使用 http.Dir 调用 wd 获取文件文件服务器相对于当前路径的相对路径
	// http://localhost:8080/static/
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{} //定义一个 handler

func (this *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //要求 handle 实现一个方法 ServeHTTP
	//io.WriteString(w, "Hello world, this is v2.")
	io.WriteString(w, "URL: "+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is v2.")
}
