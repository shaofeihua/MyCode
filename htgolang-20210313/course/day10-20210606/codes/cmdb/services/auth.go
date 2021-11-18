package services

import (
	"htgolang-20210313/course/day10-20210606/codes/cmdb/models"
)

func ValidateLogin(name string, password string) *models.User {
	if name == "kk" && password == "123456" {
		return &models.User{1, "kk", true}
	}
	return nil
}
