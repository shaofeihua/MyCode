/*
使用 beego 框架在浏览器界面输出 Hello world 。
*/
package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct { //自定义一个结构，目的是调用 beego 中已经实现的方法
	beego.Controller
}

func (this *HomeController) Get() { //Get 是 http 的方法
	this.Ctx.WriteString("Hello world.")
}

func main() {
	beego.Router("/", &HomeController{}) //注册路由
	beego.Run()
}
