package services

import (
	"cmdb/models"
)

var users = []models.User{
	models.User{1, "春景", true},
	models.User{2, "小楠", true},
	models.User{3, "莹莹", true},
	models.User{4, "混混", false},
}

func QueryUsers(q string) []models.User {
	return users
}

func DeleteUser(id int) {
	/*
		根据 ID 删除用户的逻辑：
			1、创建一个临时的空切片。
			2、遍历用户，将所有用户加入到临时的空切片。
			3、当 ID 等于用户要删除的 ID 时，跳过，该用户不加入到这个空切片中。
			4、遍历完全结束后，空切片变成 -> 除了删除的用户以外，包含其他全部的用户
			5、将这个临时的切片，赋值给原来的用户切片
	*/
	tempUsers := make([]models.User, 0, len(users))
	for _, user := range users {
		if user.ID == id {
			continue
		}
		tempUsers = append(tempUsers, user)
	}
	users = tempUsers
}
