package models

import "time"

type User struct {
	Id           int64
	Name         string
	Password     string
	Birthday     *time.Time
	Telephone    string
	Email        string
	Addr         string
	Status       int8
	RoleId       int64
	DepartmentId int64
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
	Description  string
	Sex          bool
}

// 提供给 html 模板调用的方法，用于判断性别
func (u *User) SexText() string {
	if u.Sex {
		return "男"
	}
	return "女"
}
