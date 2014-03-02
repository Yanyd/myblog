package controllers

import (
	//"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
	"path"
	"strconv"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["Title"] = "档案 • 我的博客 "
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = utils.IsLogin(this.Controller)
	if utils.IsLogin(this.Controller) {
		this.Data["Account"] = beego.AppConfig.String("account")
	}
	//管理员后台 todo
	page := utils.NewPage()
	pageNow := this.Input().Get("pageNow")
	if 0 < len(pageNow) {
		page.PageNow, _ = strconv.ParseInt(pageNow, 10, 64)
	}
	page, err := models.GetAllTopic("", "", false, page) //第一个"" 按分类显示? 第二个"" 按标签显示?
	if nil != err {
		beego.Error(err)
	} else {
		for k, topic := range page.PageContent.([]*models.Topic) {
			content := beego.Substr(beego.Html2str(topic.Content), 0, 60)
			page.PageContent.([]*models.Topic)[k].Content = content
		}
		this.Data["Page"] = page
		this.Data["Router"] = "topic"
	}

	//博客档案分页
	/*	page2 := utils.NewPage()
		pageNow2 := this.Ctx.Input.Param("1")
		if 0 < len(pageNow2) {
			page2.PageNow, _ = strconv.ParseInt(pageNow2, 10, 64)
		}
		page2.PageSize = 10

		page2, err := models.GetArchive(page2)
		if nil != err {
			beego.Error(err)
		}
		this.Data["Archive"] = archive*/
	drafts, _ := models.GetDraft()
	this.Data["Drafts"] = drafts
	this.Data["DraftNum"] = len(drafts)
	if "manage" == this.Input().Get("op") {
		showDraft := ("true" == this.Input().Get("draft"))
		if showDraft {
			this.Data["ShowDraft"] = true
		}
		this.TplNames = "topic.html"
	} else {
		this.TplNames = "topic2.html"
	}
}

func (this *TopicController) Add() {

	utils.Certification(this.Controller)

	this.Data["Title"] = "发布文章 • 我的博客 "
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = true
	this.Data["Account"] = beego.AppConfig.String("account")
	this.Data["Categories"], _ = models.GetAllCategory()
	this.TplNames = "topicAdd.html"
}

func (this *TopicController) Post() {
	utils.Certification(this.Controller)

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	cid := this.GetStrings("cid")
	tempLabel := this.Input().Get("label")

	//是否保存为草稿
	draft := ("draft" == this.Ctx.Input.Param("0"))
	//新增id=0
	id := this.Input().Get("id")
	isUpdate := ("0" != id)

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
	}
	label := []string{}
	var topicId int64
	if 0 < len(tempLabel) {
		label = strings.Split(tempLabel, ",")
	}
	if isUpdate {
		//修改草稿还是修改已经发布的博客
		err = models.UpdateTopic(id, title, content, attachmentName, cid, label, draft)
	} else {
		//新增草稿还是新增博客
		topicId, err = models.AddTopic(title, content, attachmentName, cid, label, draft)
		id = fmt.Sprintf("%d", topicId)
	}
	if nil != err {
		beego.Error(err)
	}
	//新增草稿要返回id
	this.Ctx.WriteString(`{"id":"` + id + `"}`)
	this.ServeJson()
}

func (this *TopicController) View() {
	this.Data["IsTopic"] = true
	///topic/view/123/343 -->0:123 1:343
	id := this.Ctx.Input.Param("0")
	topic, err := models.GetTopic(id)
	if nil != err {
		beego.Error(err)
		this.Redirect("/", 301)
		return
	}
	this.Data["Title"] = topic.Title
	this.Data["IsLogin"] = utils.IsLogin(this.Controller)
	if utils.IsLogin(this.Controller) {
		this.Data["Account"] = beego.AppConfig.String("account")
	}

	this.TplNames = "topicView.html"

	//注意先后顺序,不要影响文章的显示
	comments, err := models.GetAllComment(id)
	if nil != err {
		beego.Error(err)
		return
	}
	this.Data["Comments"] = comments
	label := []string{}
	if 0 < len(topic.Label) {
		label = strings.Split(topic.Label, ",")
	}
	this.Data["Labels"] = label
	this.Data["Ctgs"] = topic.TempStr
	haveAttach := (len(topic.Attachment) > 0)
	this.Data["HaveAttach"] = haveAttach
	if haveAttach {
		attachName := topic.Attachment
		if 0 < len(attachName) {
			tempType := attachName[strings.LastIndex(attachName, ".")+1:]
			if strings.Contains("jpg,png,jpeg", tempType) {
				topic.IsImage = true
			} else {
				topic.IsFile = true
			}
		}
	}
	this.Data["Topic"] = topic
}

func (this *TopicController) Update() {
	this.Data["IsTopic"] = true
	utils.Certification(this.Controller)

	id := this.Ctx.Input.Param("0")
	pageNow := this.Ctx.Input.Param("1")
	topic, err := models.GetTopic(id)
	if nil != err {
		beego.Error(err)
		this.Redirect("/", 301)
		return
	}
	this.Data["Title"] = "修改文章　" + topic.Title
	this.Data["IsLogin"] = true
	this.Data["Account"] = beego.AppConfig.String("account")
	this.Data["Topic"] = topic

	allCtg, _ := models.GetAllCategory()
	topicCtg := models.StrToSlice(topic.Cid)
	for _, category := range allCtg {
		for _, v := range topicCtg {
			tempId, _ := strconv.ParseInt(v, 10, 64)
			if tempId == category.Id {
				category.Checked = true
				break
			}
		}
	}
	this.Data["Categories"] = allCtg
	haveAttach := (len(topic.Attachment) > 0)
	this.Data["HaveAttach"] = haveAttach
	this.Data["UpdTopic"] = true
	if 0 < len(pageNow) {
		this.Data["PageNow"] = pageNow
	}
	this.TplNames = "topicUpdate.html"
}

func (this *TopicController) Del() {
	//权限验证
	utils.Certification(this.Controller)

	id := this.Ctx.Input.Param("0")
	err := models.DeleteTopic(id)
	if nil != err {
		beego.Error(err)
	}
	this.Redirect("/topic", 301)
	return
}

//用来记录点赞的IP
var ipMap = make(map[string][]string)

func (this *TopicController) Upvote() {
	id := this.Ctx.Input.Param("0")
	ip := this.Ctx.Input.IP()
	arr, ok := ipMap[ip]
	flag := false //是否重复点赞
	if ok {       //存在该Ip
		for _, value := range arr {
			if id == value { //已经点过赞了
				flag = true
				break
			}
		}
	}
	if !flag {
		upvote, err := models.Upvote(id)
		if nil != err {
			beego.Error(err)
		}
		ipMap[ip] = append(arr, id)
		this.Ctx.WriteString("{num:'" + fmt.Sprintf("%d", upvote) + "'}")
	}
	this.ServeJson()
}

func (this *TopicController) Archive() {
	date := this.Ctx.Input.Param("0")

	page := utils.NewPage()
	pageNow := this.Ctx.Input.Param("1")
	if 0 < len(pageNow) {
		page.PageNow, _ = strconv.ParseInt(pageNow, 10, 64)
	}
	page.PageSize = 6

	page, err := models.GetTopicByDate(date, page)
	if nil != err {
		beego.Error(err)
	}
	jsonObj := "{'topics':["
	length := len(page.PageContent.([]*models.Topic))
	for k, topic := range page.PageContent.([]*models.Topic) {
		id := fmt.Sprintf("%d", topic.Id)
		title := topic.Title
		content := beego.Substr(beego.Html2str(topic.Content), 0, 50)
		content = strings.Replace(content, "\\", "\\\\", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "'", "\"", -1)
		created := fmt.Sprintf("%v", topic.Created.Format("2006-01-02"))
		views := fmt.Sprintf("%d", topic.Views)
		replyCount := fmt.Sprintf("%d", topic.ReplyCount)
		label := strings.Join(models.StrToSlice(topic.Label), ",")
		ctg := topic.Cid
		cid := topic.TempStr.(string)
		upvote := fmt.Sprintf("%d", topic.Upvote)
		if k == (length - 1) {
			jsonObj += "{'id':'" + id + "','title':'" + title + "','content':'" + content + "','created':'" + created + "','views':'" + views + "','replyCount':'" + replyCount + "','label':'" + label + "','ctg':'" + ctg + "','cid':'" + cid + "','upvote':'" + upvote + "'}]"
		} else {
			jsonObj += "{'id':'" + id + "','title':'" + title + "','content':'" + content + "','created':'" + created + "','views':'" + views + "','replyCount':'" + replyCount + "','label':'" + label + "','ctg':'" + ctg + "','cid':'" + cid + "','upvote':'" + upvote + "'},"
		}
	}
	//page 对象
	jsonObj += ",'page':{'totalPageNum':'" + fmt.Sprintf("%d", page.TotalPageNum) + "','pageNow':'" + fmt.Sprintf("%d", page.PageNow) + "'}"
	jsonObj += "}"
	this.Ctx.WriteString(jsonObj)
	this.ServeJson()
}

func (this *TopicController) ArchiveDate() {
	page := utils.NewPage()
	pageNow := this.Ctx.Input.Param("0")
	if 0 < len(pageNow) {
		page.PageNow, _ = strconv.ParseInt(pageNow, 10, 64)
	}
	page.PageSize = 10
	page, err := models.GetArchive(page)
	if nil != err {
		beego.Error(err)
	}
	jsonObj := "{'dates':["
	length := len(page.PageContent.([]string))
	//在页面中js是倒排的,所以先在这边排好序
	for k := length - 1; k >= 0; k-- {
		time := page.PageContent.([]string)[k]
		if k == 0 {
			jsonObj += "{'d':'" + time + "'}]"
		} else {
			jsonObj += "{'d':'" + time + "'},"
		}
	}
	jsonObj += ",'page':{'totalPageNum':'" + fmt.Sprintf("%d", page.TotalPageNum) + "','pageNow':'" + fmt.Sprintf("%d", page.PageNow) + "'}}"
	this.Ctx.WriteString(jsonObj)
	this.ServeJson()
}
func (this *TopicController) BatchDelete() {
	//权限验证
	utils.Certification(this.Controller)

	tids := this.Ctx.Input.Param("0")
	totalPageNum, err := models.BatchDeleteTopic(strings.Split(tids, ","), utils.NewPage().PageSize)
	if nil != err {
		beego.Error(err)
	}
	this.Ctx.WriteString(`{"totalPageNum":"` + fmt.Sprintf("%d", totalPageNum) + `"}`)
	this.ServeJson()
}
