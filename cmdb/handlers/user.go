package handlers

import (
	"cmdb/services"
	"cmdb/utils"
	"cmdb/web"
	"fmt"
	"net/http"
	"strconv"
)

func QueryUsersHandler(w http.ResponseWriter, r *http.Request) {
	// 客户端发请求查询用户信息时，通常先验证是否已登录
	session := web.LoadSession(w, r)
	if _, ok := session.Get("uid").(int); !ok {
		fmt.Println("check: ", session, session.Get("uid"), session.Get("uid").(int))
		//	未登录，则跳转到登录页面
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	users := services.QueryUsers("")
	/*
		// 加载模板&执行模板
		tpl, err := template.ParseFiles("D:\\GoProject\\Mage\\cmdb\\views\\user\\users.html")
		if err != nil {
			log.Fatal(err)
		}
		err = tpl.ExecuteTemplate(w, "users.html", nil)
		if err != nil {
			log.Fatal(err)
		}
	*/

	// 使用自定义函数来加载模板&执行模板
	utils.ParseTemplate(w, r, []string{"D:\\GoProject\\src\\Mage\\cmdb\\views\\user\\users.html"}, "users.html", users)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// 要删除某用户信息时，通常先验证是否已登录
	session := web.LoadSession(w, r)
	if _, ok := session.Get("uid").(int); !ok {
		//	未登录，则跳转到登录页面
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	// 1、根据 ID 删除用户
	// 2、ID 来自客户端浏览器提交的 form
	// 3、表单中的 ID 类型为 string ，需要转换为 int
	// 4、删除完用户之后需要重新返回用户列表，即重新向 server 端请求一次用户信息
	if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
		services.DeleteUser(id)
	}
	http.Redirect(w, r, "/users/", http.StatusFound)
}
