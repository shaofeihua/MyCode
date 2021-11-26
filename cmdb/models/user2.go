package models

type User2 struct {
	ID   int
	Name string
	Sex  bool
}

// 提供给 html 模板调用的方法，用于判断性别
func (u *User2) SexText() string {
	if u.Sex {
		return "女"
	}
	return "男"
}
