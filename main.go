package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/controllers"
	"myblog/models"
	"os"
)

//进行数据库的初始化
func init() {
	models.RegisterDB()
}

//模板函数
func add(num int64) int64 {
	return num + 1
}
func sub(num int64) int64 {
	return num - 1
}

func main() {

	orm.Debug = true //打印信息
	//自动创建
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/comment", &controllers.CommentController{})
	beego.Router("/comment/add", &controllers.CommentController{}, "post:Add")
	beego.Router("/photo", &controllers.PhotoController{})
	beego.Router("/board", &controllers.BoardController{})
	beego.Router("/todo", &controllers.TodoController{})
	beego.Router("/resume", &controllers.ResumeController{})

	//使用 /comment/del/10/12 访问不到 要么使用/comment/del?cid=10&tid=12 要么使用自动路由
	//beego.Router("/comment/del/*", &controllers.CommentController{}, "get:Del")

	//自动路由 控制器+方法 剩余的解析成参数,保存在 this.Ctx.Input.Param 当中
	beego.AutoRouter(&controllers.TopicController{})
	beego.AutoRouter(&controllers.CommentController{})
	beego.AutoRouter(&controllers.TodoController{})
	beego.AutoRouter(&controllers.ResumeController{})
	beego.AutoRouter(&controllers.PhotoController{})

	//创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	//作为静态文件直接下载 在IE下会直接打开
	//beego.SetStaticPath("/attachment", "attachment") //相对路径

	//使用独立的控制器来处理静态文件,这样在IE中才会下载
	//注意: 在:all前面记得加上/
	beego.Router("/attachment/:all", &controllers.AttachController{})
	beego.Router("/findeditor", &controllers.FindEditorController{})

	beego.AddFuncMap("add", add)
	beego.AddFuncMap("sub", sub)

	beego.EnableAdmin = true

	beego.Run()
}
