package controllers

import (
	"github.com/astaxie/beego"
	//"io"
	"path"
)

type FindEditorController struct {
	beego.Controller
}

func (this *FindEditorController) Post() {
	//对中文进行解码，将url中的中文编码，解码成正常的UTF-8格式
	//处理附件
	_, fileHeader, err := this.GetFile("imgFile")
	if nil != err {
		beego.Error(err)
		return
	}
	var imgName string
	if nil != fileHeader {
		imgName = fileHeader.Filename
		beego.Info(imgName)
		//文件夹都要自己创建
		err = this.SaveToFile("imgFile", path.Join("static/kindeditor/attached/", imgName)) //使用相对路径
		//path.Join 根据系统自动添加分隔符
		if nil != err {
			beego.Error(err)
			return
		}
	}
	//他妈的，还得用双引号
	this.Ctx.WriteString("{\"error\":0,\"url\":\"/static/kindeditor/attached/" + imgName + "\"}")
	this.ServeJson()

	//this.Ctx.ResponseWriter.Write("{url:'static/kindeditor/attached/" + imgName + "'}")
}
