package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context" //新版本 beego 的 api 找不到 beego.Context 。原来的 beego.Context 改为 *context.Context
	//"net/url"
)

type LoginController struct {
	beego.Controller
}

//登录页面是 get 请求，所以这里要实现 Get() 方法
func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("username", "", -1, "/") // -1 的含义是退出后删除 cookie
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 301) //返回主页
		return                  //退出后不需要再渲染 login 页面，所以要增加一个 return
	}

	this.TplName = "login.html" //定义一个登录的模板，里面包含输入账号和密码的操作
}

//使用 post 登录，所以这里要实现 Post() 方法
func (this *LoginController) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	//return
	//获取表单信息
	username := this.Input().Get("username")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on" //直接进行判断，如果已经登录（on），则得到值“true”，反之得到“false”

	//增加判断，验证用户名和密码是否正确
	if beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0 //给 cookie 设置默认存储时限，即当浏览器关闭时 cookie 就被删除了
		if autoLogin {
			maxAge = 1<<31 - 1 //当判断为自动登录时，给 cookie 设置一个较长的存储时限
		}

		this.Ctx.SetCookie("username", username, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}

	//登陆之后重定向到首页
	this.Redirect("/", 301)
	return
}

//判断用户是否成功登陆：内容跟上面的 Post() 类似，提取 username 和 pwd 进行比对，返回 true（登陆成功） 或者 false（登陆失败）
func checkAccount(ctx *context.Context) bool {
	//获取 cookie
	ck, err := ctx.Request.Cookie("username")
	if err != nil { // err 不为 nil 的话，说明 cookie 不存在。
		return false
	}

	username := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value

	return beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("pwd") == pwd
}
