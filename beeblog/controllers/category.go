package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	//"time"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	op := this.Input().Get("op") //获取 option

	switch op { // 使用 switch 对 op 的值进行检查
	case "add": //如果 op 是添加分类的操作
		name := this.Input().Get("name") //获取用户输入的名称 name
		if len(name) == 0 {              //如果输入的 name 长度为 0 ，则跳出 switch
			break
		}

		//返回现在时间
		//tNow := time.Now()
		//时间转化为string，layout必须为 "2006-01-02 15:04:05"
		//timeNow := tNow.Format("2006-01-02 15:04:05")

		//如果错误不为空则打印出来
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		//无论是否插入成功，都要重定向，刷新获取分类列表
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id") //获取用户输入的名称 name 。删除分类通过 id
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	}
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
