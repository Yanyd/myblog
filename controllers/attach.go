package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type AttachController struct {
	beego.Controller
}

func (this *AttachController) Get() {
	//对中文进行解码，将url中的中文编码，解码成正常的UTF-8格式
	filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:]) //去掉 / 使用相对路径
	if nil != err {
		this.Ctx.WriteString(err.Error())
		return
	}
	f, err := os.Open(filePath)
	if nil != err {
		this.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()

	_, err = io.Copy(this.Ctx.ResponseWriter, f)
	if nil != err {
		this.Ctx.WriteString(err.Error())
		return
	}
}
