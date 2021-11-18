package utils

import (
	"html/template"
	"net/http"
)

// 用来解析 HTML 模板的函数
func ParseTemplate(w http.ResponseWriter, r *http.Request, files []string, tplName string, data interface{})  {
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		//return err
		panic(err)
	}
	err = tpl.ExecuteTemplate(w, tplName, data)
	if err != nil {
		//return err
		panic(err)
	}
}
