package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
	"path"
)

type PhotoController struct {
	beego.Controller
}

func (this *PhotoController) Get() {
	this.Data["Title"] = "相册 • 我的博客 "
	this.Data["IsPhoto"] = true
	this.Data["IsLogin"] = utils.IsLogin(this.Controller)
	if utils.IsLogin(this.Controller) {
		this.Data["Account"] = beego.AppConfig.String("account")
	}
	albums, err := models.GetAlbum()
	if nil != err {
		beego.Error(err)
	}
	this.Data["Albums"] = albums
	this.TplNames = "photo.html"
}

func (this *PhotoController) AddAlbum() {
	name := this.Input().Get("name")
	description := this.Input().Get("description")
	jsonStr, err := models.AddAlbum(name, description)
	if nil != err {
		beego.Error(err)
	}
	this.Ctx.WriteString(jsonStr)
	this.ServeJson()
}

func (this *PhotoController) AddPhoto() {
	aid := this.Input().Get("aid")
	description := this.Input().Get("description")
	//处理附件
	_, fileHeader, err := this.GetFile("attachment")

	if nil != err { //没有上传文件
		beego.Error(err)
		beego.Info("没有上传文件.")
	}
	var attachmentName string
	if nil != fileHeader {
		attachmentName = fileHeader.Filename
		beego.Info(attachmentName)
		err = this.SaveToFile("attachment", path.Join("static/img/photo", attachmentName)) //使用相对路径
		//path.Join 根据系统自动添加分隔符
		if nil != err {
			beego.Error(err)
		}
	}
	err = models.AddPhoto(aid, attachmentName, description)
	if nil != err {
		beego.Error(err)
	}
	this.Ctx.WriteString(`{"name":"` + attachmentName + `"}`)
	this.ServeJson()
}
