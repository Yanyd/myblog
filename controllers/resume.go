package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
	"path"
)

//留言板
type ResumeController struct {
	beego.Controller
}

func (this *ResumeController) Get() {
	this.Data["Title"] = "简历 • 我的博客 "
	this.Data["IsResume"] = true
	this.Data["IsLogin"] = utils.IsLogin(this.Controller)
	if utils.IsLogin(this.Controller) {
		this.Data["Account"] = beego.AppConfig.String("account")
	}
	resumes, err := models.GetResume()
	if nil != err {
		beego.Error(err)
	}
	this.Data["Resumes"] = resumes
	wordFile, err := models.GetWordFile()
	if nil != err {
		beego.Error(err)
	}
	if nil != wordFile {
		this.Data["WordFile"] = wordFile
	}
	this.TplNames = "resume.html"
}
func (this *ResumeController) Edit() {
	utils.Certification(this.Controller)
	this.Data["Title"] = "修改简历 • 我的博客 "
	this.Data["IsResume"] = true
	this.Data["IsLogin"] = true
	this.Data["Account"] = beego.AppConfig.String("account")
	id := this.Ctx.Input.Param("0")
	resume, err := models.GetResumeById(id)
	if nil != err {
		beego.Error(err)
	}
	this.Data["Resume"] = resume
	this.TplNames = "resumeEdit.html"
}
func (this *ResumeController) Add() {
	utils.Certification(this.Controller)
	this.Data["Title"] = "添加简历 • 我的博客 "
	this.Data["IsResume"] = true
	this.Data["IsLogin"] = true
	this.Data["Account"] = beego.AppConfig.String("account")
	this.TplNames = "resumeAdd.html"
}

func (this *ResumeController) Post() {
	utils.Certification(this.Controller)
	title := this.Input().Get("title")
	color := this.Input().Get("color")
	content := this.Input().Get("content")
	rid := this.Input().Get("rid")
	var err error
	if 0 == len(rid) { //增加
		_, err = models.AddResume(title, color, content)
	} else { //修改
		err = models.EditResume(rid, title, color, content)
	}
	if nil != err {
		beego.Error(err)
	}
	this.Redirect("/resume", 301)
}

func (this *ResumeController) Del() {
	utils.Certification(this.Controller)
	id := this.Ctx.Input.Param("0")
	err := models.DeleteResume(id)
	if nil != err {
		beego.Error(err)
	}
	this.ServeJson()
}

//上传Word文件
func (this *ResumeController) UpWord() {
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
		err = this.SaveToFile("attachment", path.Join("attachment", attachmentName)) //使用相对路径
		//path.Join 根据系统自动添加分隔符
		if nil != err {
			beego.Error(err)
		}
		id, err := models.AddResume(attachmentName, "WordFile", attachmentName)
		if nil != err {
			beego.Error(err)
		}
		this.Ctx.WriteString(`{"id":"` + id + `","name":"` + attachmentName + `"}`)
	}
	this.ServeJson()
}
