/*
	注册路由
*/
package routers

import (
	"cmdb/handlers"
	"net/http"
)

func init() {
	http.HandleFunc("/login/", handlers.LoginHandler)
	http.HandleFunc("/users/", handlers.QueryUsersHandler)
	http.HandleFunc("/user/delete/", handlers.DeleteUserHandler)
	//http.HandleFunc("/user/edit/", handlers.EditUsersHandler)
}
