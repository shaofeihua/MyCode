/*
	静态 web 服务器开发：
		1、定义 url 对应的处理器/处理器函数
		2、启动 web 服务
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	/*
		实现一个简单的静态 web 服务器
	*/
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

// 处理器要实现的函数
func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取用户请求的 URL 使用参数 r
	fmt.Println("r.URL:", r.URL) // 获取 uri ，包含问号及后面的参数
	fmt.Println("r.URL.Scheme:", r.URL.Scheme)
	fmt.Println("r.URL.Opaque:", r.URL.Opaque)
	fmt.Println("r.URL.User:", r.URL.User)
	fmt.Println("r.URL.Host:", r.URL.Host)
	fmt.Println("r.URL.Path:", r.URL.Path) // 获取 uri ，不带问号及后面的参数
	fmt.Println("r.URL.RawPath:", r.URL.RawPath)
	fmt.Println("r.URL.ForceQuery:", r.URL.ForceQuery) // bool 类型。判断 url 是否带参数，即是否带问号及其问号后面是否有参数。默认为 true，即没有参数。如果为 false ，说明有参数
	fmt.Println("r.URL.RawQuery:", r.URL.RawQuery)     // 获取 url 问号后面的参数
	fmt.Println("r.URL.Fragment:", r.URL.Fragment)
	fmt.Println("r.URL.RawFragment:", r.URL.RawFragment)

	//currentPath, errPWD := os.Executable()
	//if errPWD != nil {
	//	fmt.Println("errPWD:", errPWD)
	//}
	//fmt.Printf("file 的类型是: %T\n", currentPath)
	//fmt.Println("currentPath:", currentPath)
	fmt.Println("h.dir:", h.dir)
	fmt.Println("转换前的 r.URL.Path:", r.URL.Path)

	// 将 r.URL.Path 的路径转换为 windows 下的路径
	r.URL.Path = filepath.FromSlash(r.URL.Path)
	fmt.Println("转换后的 r.URL.Path:", r.URL.Path)

	//path := h.dir + r.URL.Path
	path := "D:\\GoCode\\Mage" + r.URL.Path
	fmt.Println("file:", path)
	fmt.Printf("file 的类型是: %T\n", path)
	if fhandler, err := os.Open(path); err == nil {
		fmt.Println("fhandler:", fhandler)
		if _, errRead := io.Copy(w, fhandler); errRead != nil { // 读取文件
			//log.Fatal(errRead)
			fmt.Println("errRead:", errRead)
		}
		if errClose := fhandler.Close(); errClose != nil {
			//log.Fatal(errClose)
			fmt.Println("errClose:", errClose)
		}
	} else {
		//log.Fatal(err)
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusNotFound) // 没找到文件，返回 404
	}
}
