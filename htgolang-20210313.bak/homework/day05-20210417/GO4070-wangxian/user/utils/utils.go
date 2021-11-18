package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"
	"user/module"
)

//md5加密
func Md5Encrypt(b []byte) string {
	hasher := md5.New()
	hasher.Write(b)
	hasher.Write([]byte("asxeft"))
	hash := fmt.Sprintf("%x", hasher.Sum(nil))
	return hash
}

//获取用户输入信息
func Input(msg string) (value string) {
	fmt.Print(msg)
	fmt.Scanln(&value)
	return strings.TrimSpace(value)
}

//检查用户是否存在,存在返回在切片中的索引，不存在返回-1
func CheckUserExist(uid string, users []module.User) int {
	id, _ := strconv.Atoi(uid)
	for i, v := range users {
		if id == v.Id {
			return i
		}
	}
	return -1
}

//用户输入task信息
func InputTask(id int) module.Task {
	name := Input("任务名称:")
	priority := Input("任务优先级:")
	desc := Input("任务描述:")
	status := Input("任务状态:")
	progress := Input("任务进度:")
	owner := Input("任务负责人:")

	tmpTask := module.Task{id, name, priority, desc, status, progress, owner, time.Now(), nil}
	return tmpTask
}


//时间解析
func TimeParse(t *time.Time)  {
	
}