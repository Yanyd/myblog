package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Add() {
	//utils.Certification(this.Controller)
	id := this.Input().Get("tid")
	nickname := this.Input().Get("nickname")
	email := this.Input().Get("email")
	content := this.Input().Get("content")
	pid := this.Input().Get("pid")
	//todo:在代码层面进行邮件的验证
	/*	if !utils.EmailVerification(email) {
		this.Ctx.WriteString("{error:'1'}")
	} else {*/
	comment, err := models.AddComment(id, nickname, email, content, pid, utils.NewPage().PageSize)
	if nil != err {
		beego.Error(err)
	}
	this.Ctx.WriteString(comment)
	//}
	this.ServeJson()
}

func (this *CommentController) Del() {
	utils.Certification(this.Controller)
	commentId := this.Ctx.Input.Param("0")
	topicId := this.Ctx.Input.Param("1")
	jsonStr, err := models.DeleteComment(commentId, topicId, utils.NewPage().PageSize)
	if nil != err {
		beego.Error(err)
	}
	this.Ctx.WriteString(jsonStr)
	this.ServeJson()
	/*	this.Redirect("/topic/view/"+topicId, 301)
		return*/
}

func (this *CommentController) GetNew() {
	//多余？直接写8?
	page := utils.NewPage()
	page.PageSize = 8
	jsonStr, err := models.GetNewComment(page.PageSize)
	if nil != err {
		beego.Error(err)
	}
	this.Ctx.WriteString(jsonStr)
	this.ServeJson()
}
