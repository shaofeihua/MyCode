package handlers

import (
	"htgolang-20210313/course/day10-20210606/codes/cmdb/services"
	"htgolang-20210313/course/day10-20210606/codes/cmdb/utils"
	"htgolang-20210313/course/day10-20210606/codes/cmdb/web"
	"net/http"
	"strconv"
)

func QueryUserHandler(w http.ResponseWriter, r *http.Request) {
	// 验证用户是否登陆
	token := web.LoadToken(w, r) // 加载token
	if _, ok := token.Get("uid").(int); !ok {
		// 已登录, 跳转到登陆页面
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	users := services.QueryUsers("")
	utils.ParseTemplate(w, r, []string{"D:/GoProject/src/htgolang-20210313/course/day10-20210606/codes/cmdb/views/user/users.html"}, "users.html", users)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// 验证用户是否登陆
	token := web.LoadToken(w, r) // 加载token
	if _, ok := token.Get("uid").(int); !ok {
		// 已登录, 跳转到登陆页面
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
		services.DeleteUser(id)
	}
	http.Redirect(w, r, "/users/", http.StatusFound)

}
