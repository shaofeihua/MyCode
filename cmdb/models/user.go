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

func (u *User) SexText() string {
	if u.Sex {
		return "男"
	} else {
		return "女"
	}
}
