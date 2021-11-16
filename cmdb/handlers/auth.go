package handlers

import (
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session := web.LoadSession(w, r) // 加载 session

	err := ""
	name := ""
	password := ""
	// 判断请求的方法
	if r.Method == http.MethodPost {

		// 登录验证：判断用户名和密码是否正确。用户名和密码来自客户端浏览器提交的表单中的 form
		name = r.PostFormValue("name")
		password = r.PostFormValue("password")

		// 将用户输入的用户名和密码打印到日志中
		fmt.Println(name, password)

		if user := services.ValidataLogin(name, password); user != nil {
			// 用户名和密码验证成功，登录成功后跳转到其他位置
			session.Set("uid", user.ID) // 设置用户登录状态。框架中没有，需要自己写
			web.DumpSession(w,r,session)    // 将 session 持久化存储。框架中通常已有，不需要自己写
			http.Redirect(w, r, "/users/", http.StatusFound)
			return // 这里为什么要返回？？？
		}
		// 如果用户名和密码验证失败，要返回登录页面，并且把用户提交的用户名回填进去。这个回填用户名的功能在 login.html 模板中实现
		err = "用户名或者密码错误，请重新输入"
	}

	/*
		// 加载模板&执行模板
		tpl, err := template.ParseFiles("D:\\GoProject\\Mage\\cmdb\\views\\auth\\login.html")
		if err != nil {
			// panic(err)
			log.Fatal(err)
		}
		err = tpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			// panic(err)
			log.Fatal(err)
		}
	*/

	// 使用自定义函数来加载模板&执行模板
	//utils.ParseTemplate(w, r, []string{"D:\\GoProject\\src\\Mage\\cmdb\\views\\auth\\login.html"}, "login.html", struct {
	utils.ParseTemplate(w, r, []string{"D:/GoProject/src/Mage/cmdb/views/auth/login.html"}, "login.html", struct {
		Name string
		Err  string
	}{name, err})
}
