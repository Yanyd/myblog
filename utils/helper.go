package utils

import (
	"github.com/astaxie/beego"
	"net/smtp"
)

func IsLogin(this beego.Controller) bool {
	account := this.GetSession("account")
	if nil == account {
		return false
	}
	account = account.(string)
	pwd := this.GetSession("pwd")
	if nil == pwd {
		return false
	}
	pwd = pwd.(string)
	return beego.AppConfig.String("account") == account && beego.AppConfig.String("pwd") == pwd
}

func Certification(this beego.Controller) {
	//权限验证
	if !IsLogin(this) {
		this.Abort("401")
	}
}

func EmailVerification(email string) bool {
	/*	//开启权限验证
		auth := smtp.PlainAuth("", "yanyuanda@sina.cn", "", "snmpt.sina.com")
		//发送消息
		err := smtp.SendMail("snmpt.sina.com:25", auth, "yanyuanda@sina.cn", []string{email,"yanyuanda@si"}, []byte("Subject:hello\r\n\r\n12312"))
		//单个接收邮件不会发送成功(不会100%成功)!!! 被360拦截?
		if nil != err {
			beego.Error(err)
			println("-----------------------------------")
			return false
		}
		return true*/

	client, err := smtp.Dial("smtp.sina.com:25")
	if nil != err {
		beego.Error(err)
		return false
	}
	auth := smtp.PlainAuth("", "yanyuanda@sina.cn", "", "smtp.sina.com")
	err = client.Auth(auth)
	if nil != err {
		beego.Error(err)
		return false
	}
	err = client.Verify(email)
	if nil != err {
		beego.Error(err)
		return false
	}
	return true
}
