package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err.Error)
	} else {
		this.Data["Topics"] = topics
	}

}

func (this *TopicController) Post() {
	if checkAccount(this.Ctx) { //添加文章之前先验证用户身份（博客主人才能添加）
		this.Redirect("/login", 301) //如果身份验证未通过，则跳转到登录页面
		return
	}

	//对用提交的表单进行解析
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")

	var err error
	//判断：如果 tid 等于 0，则是新增文章。如果不等于 0 ，则是修改文章
	if len(tid) == 0 {
		err = models.AddTopic(title, content, category)
	} else {
		err = models.ModifyTopic(title, content, category)
	}
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 301)
}

//使用自动路由
func (this *TopicController) Add() {
	//this.Ctx.WriteString("add")
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	//topic, err := models.GetTopic(this.Ctx.Input.Params())
	topic, err := models.GetTopic(this.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 301)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = this.Ctx.Input.Params()
}

func (this *TopicController) Modify() {
	this.TplName = "Topic_modify.html"

	//修改文章之前先获取文章
	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 301) //获取失败则跳转到首页
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}

func (this *TopicController) Delete() {
	//先验证是否登录。如果未登录则跳转到登录界面
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 301)
		return
	}
	err := models.DelTopic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/", 301)
}
