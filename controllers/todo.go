package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
)

//待办事项
type TodoController struct {
	beego.Controller
}

func (this *TodoController) Get() {
	this.Data["Title"] = "Todo • 我的博客 "
	this.Data["IsTodo"] = true
	this.Data["IsLogin"] = utils.IsLogin(this.Controller)
	if utils.IsLogin(this.Controller) {
		this.Data["Account"] = beego.AppConfig.String("account")
	}
	todos, err := models.GetAllTodo()
	if nil != err {
		beego.Error(err)
	}
	this.Data["Todos"] = todos
	this.TplNames = "todo.html"
}

func (this *TodoController) Post() {
	content := this.Input().Get("content")
	id, err := models.AddTodo(content)
	if nil == err {
		this.Ctx.WriteString(`{"ok":true,"id":"` + id + `"}`)
	} else {
		beego.Error(err)
	}
	this.ServeJson()
}

func (this *TodoController) UpdStatus() {
	id := this.Ctx.Input.Param("0")
	err := models.UpdateStatus(id)
	if nil != err {
		beego.Error(err)
	} else {
		this.Ctx.WriteString(`{"ok":true}`)
	}
	this.ServeJson()
}
