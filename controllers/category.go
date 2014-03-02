package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {

	op := this.Input().Get("op")
	switch op {
	case "add":
		utils.Certification(this.Controller)
		name := this.Input().Get("name")
		if 0 == len(name) {
			break
		}
		err := models.AddCategory(name)
		if nil != err {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		//权限验证
		utils.Certification(this.Controller)
		cid := this.Input().Get("cid")
		if 0 == len(cid) {
			break
		}
		err := models.DelCategory(cid)
		if nil != err {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}

	this.Data["Title"] = "分类 • 我的博客 "
	this.Data["IsCategory"] = true
	this.Data["IsLogin"] = utils.IsLogin(this.Controller)
	if utils.IsLogin(this.Controller) {
		this.Data["Account"] = beego.AppConfig.String("account")
	}
	//this.Data 已是定义好了的类型,且不可改变,不可用 := 赋值
	page := utils.NewPage()
	pageNow := this.Input().Get("pageNow")
	if 0 < len(pageNow) {
		page.PageNow, _ = strconv.ParseInt(pageNow, 10, 64)
	}

	this.Data["Page"], _ = models.GetAllCategoryByPage(page)
	this.Data["Router"] = "category"
	this.Data["AllCtg"], _ = models.GetAllCategory()
	this.TplNames = "category.html"
}
