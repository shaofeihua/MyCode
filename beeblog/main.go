package main

import (
	//"beeblog/controllers"
	"beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	//models.RegisterDB()
	orm.Debug = true                      //开启 ORM 调试模式
	orm.RunSyncdb("default", false, true) //自动建表。默认情况下不自动建表。第一个参数是只数据库名，第二个参数如果是 true，则不管表是否存在，都会删除重建。所以安全起见，使用 false，如果表已经存在则跳过。第三个参数如果是 true，就打印详细的信息

	//beego.Router("/", &controllers.MainController{})          //注册首页路由。
	//beego.Router("/login", &controllers.LoginController{})    //登录页面的路由。
	//beego.Router("/topic", &controllers.TopicController{})    //文章页面的路由。
	//beego.Router("/category", &controllers.CategoryController{}) //分类页面的路由。

	beego.Run() //启动 beego
}
