package models

type User struct {
	ID   int
	Name string
	Sex  bool
}

// 提供给 html 模板调用的方法，用于判断性别
func (u *User) SexText() string {
	if u.Sex {
		return "女"
	}
	return "男"
}
