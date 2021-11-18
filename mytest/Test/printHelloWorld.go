package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// request.FormValue 读取请求地址url中的参数。浏览器中的访问方式为：http://localhost:8080/?name=******
		fmt.Fprintf(writer, "<h1>Hello world %s!</h1>", request.FormValue("name"))
	})

	http.ListenAndServe(":8080", nil)
}
