//cookie 的存取操作
package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", Cookie) //注册路由
	htto.handleFunc("/2", Cookie2)

	http.ListenAndServe(":8080", nil)
}

func Cookie(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{ //Cookie 是一个结构体。可在 gowalker.org/net/http#Cookie 中查看
		Name:   "myCookie",
		Value:  "hello",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120, //单位是秒
	}

	http.SetCookie(w, ck)

	//读取 cookie
	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, ck2.Value)
}

//使用 handler 的方法设置 cookie
func Cookie2(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{ //Cookie 是一个结构体。可在 gowalker.org/net/http#Cookie 中查看
		Name:   "myCookie",
		Value:  "helloworld",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120, //单位是秒
	}

	w.Header().Set("Set-Cookie", ck.String())

	//读取 cookie
	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, ck2.Value)
}
