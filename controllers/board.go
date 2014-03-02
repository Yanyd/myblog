package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
	"strconv"
)

//留言板
type BoardController struct {
	beego.Controller
}

func (this *BoardController) Get() {
	this.Data["Title"] = "留言 • 我的博客 "
	this.Data["IsBoard"] = true
	this.Data["IsLogin"] = utils.IsLogin(this.Controller)
	if utils.IsLogin(this.Controller) {
		this.Data["Account"] = beego.AppConfig.String("account")
	}
	//取得所有留言
	//comments, err := models.GetAllComment("0")
	page := utils.NewPage()
	pageNow := this.Input().Get("pageNow")
	if 0 < len(pageNow) {
		page.PageNow, _ = strconv.ParseInt(pageNow, 10, 64)
	}
	page, err := models.GetCommentByPage("0", page)
	if nil != err {
		beego.Error(err)
		return
	}
	this.Data["Page"] = page
	this.Data["Router"] = "board"
	//暂时用StartIndex代替总的留言数(包含非顶级留言)
	this.Data["CommentNum"] = page.StartIndex

	this.TplNames = "board.html"
}
