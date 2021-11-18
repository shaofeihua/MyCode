package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})             //注册首页路由。
	beego.Router("/login", &controllers.LoginController{})       //登录页面的路由。
	beego.Router("/category", &controllers.CategoryController{}) //分类页面的路由。
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	//beego.Router("/topic", &controllers.TopicController{})
	//beego.AutoRouter(&controllers.TopicController{})
	//beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	//beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Del")
}
