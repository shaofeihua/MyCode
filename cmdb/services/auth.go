package services

import (
	"cmdb/models"
)

// 验证用户登录使用的用户名和密码是否正确
func ValidataLogin(name string, password string) *models.User {
	// 正常来讲应该是从数据库中查询用户名及其密码
	if name == "chunjing" && password == "123" {
		return &models.User{1, "chunjing", true}
	}
	return nil
}
