package services

import (
	"cmdb/models"
)

const sqlQueryUserByName = "select id,name,password from user where name=? limit 1"
// 验证用户登录使用的用户名和密码是否正确
func ValidataLogin(name string, password string) *models.User {
	var user models.User
	var pwd string
	//db.QueryRow("select id,name,password from user limit 1").Scan(&user.Id,&user.Name,&user.Password)
	err:=db.QueryRow(sqlQueryUserByName,name).Scan(&user.Id,&user.Name,&pwd)
	// 先用明文，主要产品或者项目中要使用 bcrypt 加密方式或者加盐 sha256
	if err == nil && password == pwd {
		return &user
	}
	return nil
}
