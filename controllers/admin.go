package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Title"] = "首页 • 我的博客 "
	this.Data["IsIndex"] = true
	this.Data["IsLogin"] = utils.IsLogin(this.Controller)
	if utils.IsLogin(this.Controller) {
		this.Data["Account"] = beego.AppConfig.String("account")
	}
	cid := this.Input().Get("id")
	label := this.Input().Get("label")

	tempMap := make(map[string]string)
	tempMap["id"] = cid
	tempMap["label"] = label
	this.Data["IndexMap"] = tempMap
	page := utils.NewPage()
	pageNow := this.Input().Get("pageNow")
	if 0 < len(pageNow) {
		page.PageNow, _ = strconv.ParseInt(pageNow, 10, 64)
	}
	page, err := models.GetAllTopic(cid, label, true, page)
	if nil != err {
		beego.Error(err)
	} else {
		//判断附件类型
		for k, topic := range page.PageContent.([]*models.Topic) {

			//现在没什么用处
			attachName := topic.Attachment
			if 0 < len(attachName) {
				tempType := attachName[strings.LastIndex(attachName, ".")+1:]
				if strings.Contains("jpg,png,jpeg", tempType) {
					topic.IsImage = true
				} else {
					topic.IsFile = true
				}
			}
			//// end

			content := beego.Substr(beego.Html2str(topic.Content), 0, 200)
			content = strings.Replace(content, "\\", "\\\\", -1)
			content = strings.Replace(content, "\n", "", -1)
			content = strings.Replace(content, "\r\n", "", -1)
			content = strings.Replace(content, "'", "\"", -1)
			page.PageContent.([]*models.Topic)[k].Content = content
		}
		this.Data["Page"] = page
	}
	this.TplNames = "index.html"
	categories, err := models.GetAllCategory()
	if nil != err {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
	this.Data["Labels"] = models.GetAllLabel()
}
