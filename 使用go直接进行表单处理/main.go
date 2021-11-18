//使用 go 直接进行表单（用户名和密码登录）处理
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hey)
	http.ListenAndServe(":8080", nil) //没有自己的 handler，所以直接传 nil
}

//将模板文件的内容集成在本文件中
const tpl = `
<html>
	<head>
		<title>Hey</title>
	</head>
	<body>
		<form method="post" action="/">
			Username: <input type ="text" name="uname">
			Password: <input type ="password" name="pwd">
			<button type="submit">Submit</button>
		</form>
	</body>
</html>
`

func Hey(w http.ResponseWriter, r *http.Request) { //ResponseWriter 是个接口; Request是个结构，需要以传地址的方式传进来
	//先进行判断，如果是 Get 方法，那么先把表单（用户名密码登录）页面显示出来
	if r.Method == "GET" {
		//进行 Get 请求的时候，需要将上面的 tpl 输出。此时需要先创建一个模板对象
		t := template.New("hey")
		t.Parse(tpl)
		t.Execute(w, nil)

	} else {
		//解析
		fmt.Println(r.FormValue("uname"))
	}
}
