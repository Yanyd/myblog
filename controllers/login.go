package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	//"github.com/astaxie/beego/session"
)

type LoginController struct {
	beego.Controller
}

/*var globalSession *session.Manager

func init() {
	globalSession, _ = session.NewManager("memory", "user", 3600, "", false, "sha1", "sessionidkey", 3600)
	go globalSession.GC()
}*/

func (this *LoginController) Get() {
	isExist := ("true" == this.Input().Get("exit"))
	if isExist {
		//清除session
		/*		this.DestroySession()
				//客户端cookie存放时间
				beego.SessionCookieLifeTime = -1
				//服务器端
				beego.SessionGCMaxLifetime = -1*/
		/*	this.SetSession("account", nil)
			this.SetSession("pwd", nil)
			println(utils.IsLogin(this.Controller))*/
		this.DestroySession()
		this.Redirect("/", 301)
		return
	} else {
		flash := beego.ReadFromRequest(&this.Controller)
		if v, ok := flash.Data["error"]; ok {
			this.Data["Error"] = v
			this.Data["HasError"] = true
		}
	}

	this.Data["Title"] = "登录 • 我的博客 "
	this.TplNames = "login.html"
}

//登录页面不能进行缓存

func (this *LoginController) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input())) //打印信息

	account := this.Input().Get("account")
	pwd := this.Input().Get("pwd")
	autoLogin := ("1" == this.Input().Get("autoLogin"))

	if beego.AppConfig.String("account") == account && beego.AppConfig.String("pwd") == pwd {
		//设置cookie
		maxAge := 0 //缓存时间设置
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		this.SetSession("account", account)
		this.SetSession("pwd", pwd)
		this.SetSession("expire", maxAge)
		this.Redirect("/", 301) //重定向到首
	} else {
		//返回登录页面并提示错误信息
		flash := beego.NewFlash()
		flash.Error("用户名或密码错误!")
		flash.Store(&this.Controller)
		this.Redirect("/login", 301)
	}

	return //注意return,避免重新渲染
}
