package main

import (
	"embed"
	"fmt"
	"html/template"
	"os"
)

//go:embed tpl
var fs embed.FS

func main() {
	login, err := fs.ReadFile("tpl/login.html")
	fmt.Println(string(login), err)

	// 方法一：
	tpl, err := template.New("").Parse(string(login))
	tpl.Execute(os.Stdout, "who am i")
	fmt.Println()

	// 方法二：
	templ, err := template.ParseFS(fs,"tpl/*.html")
	fmt.Println(err, templ)
	templ.ExecuteTemplate(os.Stdout, "login.html", "I am ironman")
	fmt.Println()
	templ, err = template.ParseFS(fs,"tpl/users/*.html")
	templ.ExecuteTemplate(os.Stdout, "add.html", "it is add")
	fmt.Println()
}
