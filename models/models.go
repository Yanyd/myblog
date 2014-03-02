package models

import (
	//"github.com/Unknwon/com" //通用函数包
	"github.com/astaxie/beego/orm"
	//_ "github.com/mattn/go-sqlite3" //数据库驱动 _ 表示只执行初始化函数,而不进行其他函数调用
	_ "github.com/go-sql-driver/mysql"
	//"os"
	//"path"
	"fmt"
	"myblog/utils"
	"os"
	"path"
	"strconv"
	"strings"
	//"sync"
	"encoding/base64" //Base64是一种基于64个可打印字符来表示二进制数据的表示方法
	"net/mail"
	"net/smtp"
	"time"
)

/*const (
	_DB_NAME        = "data/myblog.db" //数据库名称
	_SQLITE3_DRIVER = "mysql"          //驱动
)*/

//分类
type Category struct {
	Id            int64     `orm:"auto"`
	Title         string    //默认255字节
	Created       time.Time `orm:"index"` //建立索引  ``反射时使用 orm表示orm才会去读这个值
	Views         int64     `orm:"index"` //浏览数量
	TopicTime     time.Time `orm:"index"` //最后一篇文章的创建时间
	TopicCount    int64     //分类下的文章数量
	CtgLastUserId int64     //最后操作分类的用户
	Checked       bool      `orm:"-"` //修改时有用 checkbox
}

//文章
type Topic struct {
	Id              int64 `orm:"auto"`
	Uid             int64
	Title           string
	Content         string    `orm:"size(20000)"`
	Attachment      string    //附件(文件/图片)
	Created         time.Time `orm:"index"` //建立索引  ``反射时使用 orm表示orm才会去读这个值
	Updated         time.Time `orm:"index"` //更新时间
	Views           int64     `orm:"index"` //浏览数量
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
	Cid             string      //所属分类
	Label           string      //文章标签
	IsImage         bool        `orm:"-"` //附件类型 image
	IsFile          bool        `orm:"-"` //附件类型 file
	Upvote          int64       //点赞数量
	TempStr         interface{} `orm:"-"` //临时字段 添加各种值
	Draft           bool        //是否是草稿
}

//评论 和 留言(Tid=0)
type Comment struct {
	Id       int64 `orm:"auto"`
	Tid      int64
	Nickname string
	Content  string    `orm:"size(1000)"`
	Created  time.Time `orm:"index"`
	Pid      int64     //对哪条评论进行回复
	Email    string
	Floor    int64 //回复楼层
	Website  string
}

//todo 列表
type Todo struct {
	Id      int64  `orm:"auto"`
	Content string `orm:"size(1000)"`
	Start   time.Time
	Done    bool
}

//album 相册
type Album struct {
	Id          int64 `orm:"auto"`
	Name        string
	Description string
	Created     time.Time
	Cover       string `orm:"-"` //封面
	Num         int64  `orm:"-"` //图片数量
}

//photo 图片
type Photo struct {
	Id          int64 `orm:"auto"`
	AlbumId     int64 //属于哪个相册
	Name        string
	Description string //简述
	Created     time.Time
}

//resume 简历 按模块编写
type Resume struct {
	Id      int64  `orm:"auto"`
	Title   string //模块标题
	Color   string //背景颜色,用于标识某个模块的重要程度  alert-info/danger/success
	Content string `orm:"size(10000)"`
}

type Admin struct {
	Id        int64 `orm:"auto"`
	Name      string
	Email     string
	LoginPwd  string
	EmailHost string
	EmailPwd  string
}

//建立完模型后,记得在数据库中注册模型,然后重启服务器,因为init只执行一次

//使用orm进行数据库创建
//由主程序main进行调用
func RegisterDB() {
	//首先判断数据库文件存在与否
	/*	if !com.IsExist(_DB_NAME) {
		//需要手动创建目录  path.Dir获取路径中的目录
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		//创建数据库文件
		os.Create(_DB_NAME)
	}*/
	orm.RegisterModel(new(Category), new(Topic), new(Comment), new(Todo), new(Album), new(Photo), new(Resume), new(Admin))
	//orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_MySQL)
	//orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10) // 10最大连接数
	orm.RegisterDataBase("default", "mysql", "root:root@/myblog_db?charset=utf8", 10) // 10最大连接数
	//初始化类别表 自动添加杂七杂八类别
	AddCategory("杂七杂八")

}

/*        分类相关操作开始     */
func AddCategory(name string) error {
	o := orm.NewOrm()
	ctg := &Category{Title: name, Created: time.Now()}
	queryTable := o.QueryTable("category")
	err := queryTable.Filter("title", name).One(ctg)
	if nil == err { //已经存在
		if "杂七杂八" == name {
			fmt.Println("已经添加'杂七杂八'类别.")
		} else {
			fmt.Println("类别'" + name + "'已经存在了.")
		}
		return err
	}
	_, err = o.Insert(ctg)
	if nil != err { //插入失败
		return err
	}
	if "杂七杂八" == name {
		fmt.Println("已经添加'杂七杂八'类别.")
	}
	return nil
}

func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	queryTable := o.QueryTable("category").OrderBy("-created")
	_, err := queryTable.All(&categories)
	return categories, err
}

func GetAllCategoryByPage(page *utils.Page) (*utils.Page, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	qt := o.QueryTable("category").OrderBy("-created")
	var err error
	page.TotalRecNum, err = qt.Count()
	_, err = qt.Limit(page.PageSize, (page.PageNow-1)*page.PageSize).All(&categories)
	page.PageContent = categories

	page.TotalPageNum = (page.TotalRecNum-1)/page.PageSize + 1
	page.PrevPage = page.PageNow > 1
	page.NextPage = page.PageNow < page.TotalPageNum

	return page, err
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64) //10进制,64位
	if nil != err {
		return err
	}
	o := orm.NewOrm()
	ctg := &Category{Id: cid}
	_, err = o.Delete(ctg)
	if nil != err {
		return err
	}
	//删除属于此分类的文章的分类值
	topics := make([]*Topic, 0)
	_, err = o.QueryTable("topic").Filter("cid__contains", "$"+id+"#").All(&topics)
	if nil != err {
		return err
	}
	for _, t := range topics {
		topic := &Topic{Id: t.Id}
		err = o.Read(topic)
		if nil == err {
			topic.Cid = strings.Replace(topic.Cid, "$"+id+"#", "", -1)
			o.Update(topic)
		}
	}
	return err
}

/*       分类相关操作结束      */

/*       博客相关操作开始*/
func AddTopic(title, content, attachmentName string, cid, label []string, draft bool) (int64, error) {
	tLabel := ""
	if 0 < len(label) {
		tLabel = "$" + strings.Join(label, "#$") + "#"
	}
	if 0 == len(cid) {
		cid = []string{"1"}
	}
	tCid := "$" + strings.Join(cid, "#$") + "#"
	//todo: 分类和标签都要判断是否为空
	o := orm.NewOrm()
	topic := &Topic{
		Title:      title,
		Content:    content,
		Attachment: attachmentName,
		Created:    time.Now(),
		Updated:    time.Now(),
		Cid:        tCid,
		Label:      tLabel,
		Draft:      draft,
	}
	_, err := o.Insert(topic)
	if nil != err { //插入失败
		return 0, err
	}
	//相应的分类的文章数也要增加 如果是草稿,就不要加了，等到发布后才加
	if !draft {
		for _, id := range cid {
			ctgId, err := strconv.ParseInt(id, 10, 64)
			if nil != err {
				return 0, err
			}
			o := orm.NewOrm()
			ctg := &Category{Id: ctgId}
			err = o.Read(ctg)
			if nil != err {
				return 0, err
			} else { //查找到了
				ctg.TopicCount++
				o.Update(ctg)
			}
		}
	}
	return topic.Id, err
}

//取出草稿
func GetDraft() (topics []*Topic, err error) {
	o := orm.NewOrm()
	qt := o.QueryTable("topic")
	topics = make([]*Topic, 0)
	_, err = qt.Filter("draft", true).All(&topics, "id", "title", "updated")
	if nil != err {
		return
	}
	/*	for k, topic := range topics {
		//topics[k].Label = strings.Join(StrToSlice(topic.Label), ",")
	}*/
	return
}

func GetAllTopic(id, label string, isDesc bool, page *utils.Page) (*utils.Page, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qt := o.QueryTable("topic").Filter("draft", false)
	var err error
	if isDesc {
		//首页按更新时间倒序显示文章
		if 0 < len(id) { //按分类显示
			//文章有多个分类时怎么办 1,2,3,14, 1,--> $+ #$ +#
			qt = qt.Filter("cid__contains", "$"+id+"#")
		}
		if 0 < len(label) {
			qt = qt.Filter("label__contains", "$"+label+"#")
		}
		qt = qt.OrderBy("-updated")
	}
	page.TotalRecNum, err = qt.Count()
	_, err = qt.Limit(page.PageSize, (page.PageNow-1)*page.PageSize).All(&topics)
	page.PageContent = topics

	page.TotalPageNum = (page.TotalRecNum-1)/page.PageSize + 1
	page.PrevPage = page.PageNow > 1
	page.NextPage = page.PageNow < page.TotalPageNum

	return page, err
}

func GetTopic(id string) (topic *Topic, err error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return nil, err
	}
	o := orm.NewOrm()
	qt := o.QueryTable("topic")
	topic = new(Topic)
	err = qt.Filter("id", tid).One(topic)
	if nil != err {
		return nil, err
	}
	//浏览次数加1
	topic.Views++
	_, err = o.Update(topic)
	//复原
	topic.Label = strings.Join(StrToSlice(topic.Label), ",")
	//topic.Cid = strings.Join(, ",")
	tempMap := make(map[int64]string)
	allCtg, _ := GetAllCategory()
	for _, cid := range StrToSlice(topic.Cid) {
		tempCid, _ := strconv.ParseInt(cid, 10, 64)
		for _, ctg := range allCtg {
			if tempCid == ctg.Id {
				tempMap[ctg.Id] = ctg.Title
				break
			}
		}
	}
	topic.TempStr = tempMap
	return
}

func UpdateTopic(id, title, content, attachmentName string, cid, label []string, draft bool) (err error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tid}
	err = o.Read(topic)
	if nil != err {
		return
	} else { //查找到了
		oldCid := StrToSlice(topic.Cid)
		//1、添加时没有附件，修改时可直接覆盖 ，不用删除附件
		//2、添加时有附件，修改时没附件，不变 ，不用删除附件
		//3、添加时有附件，修改时有附件，直接覆盖，删除旧的附件
		oldAttach := topic.Attachment
		oldDraft := topic.Draft
		if 0 < len(attachmentName) {
			topic.Attachment = attachmentName
		}
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		topic.Draft = draft

		if 0 == len(cid) {
			cid = []string{"1"}
		}
		topic.Cid = "$" + strings.Join(cid, "#$") + "#"

		tLabel := ""
		if 0 < len(label) {
			tLabel = "$" + strings.Join(label, "#$") + "#"
		}
		topic.Label = tLabel

		o.Update(topic)

		//删除旧的附件
		if 0 < len(oldAttach) && 0 < len(attachmentName) {
			os.Remove(path.Join("attachment", oldAttach))
		}

		//相应的分类的文章数也需要修改，需与原来的对比
		//去掉的减1,增加的加1,没修改的不变
		//这边考虑了一篇博客对应多个分类的情况
		if !draft {
			err = checkCtg(oldCid, cid, true, oldDraft)
			err = checkCtg(cid, oldCid, false, oldDraft)
		}

		return
	}
}

//var once sync.Once

func checkCtg(cid, ids []string, isNew, oldDraft bool) error {

	for _, id := range ids {
		ctgId, err := strconv.ParseInt(id, 10, 64)
		if nil != err {
			return err
		}
		o := orm.NewOrm()
		ctg := &Category{Id: ctgId}
		err = o.Read(ctg)
		if nil != err {
			return err
		} else { //查找到了
			if isContain(id, cid) { //原来就有的
				//如果本来是草稿 就要加1 因为保存草稿时,并没有去操作类别
				if oldDraft && isNew {
					//once.Do(func() { //调用两次checkCtg会重复加
					ctg.TopicCount++
					//})
				}
			} else {
				//新增的类别
				if isNew {
					ctg.TopicCount++
				} else {
					ctg.TopicCount--
				}
			}
			o.Update(ctg)
		}
	}
	return nil
}
func isContain(elem string, arr []string) bool {
	for _, v := range arr {
		if elem == v {
			return true
		}
	}
	return false
}

func DeleteTopic(id string) (err error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return err
	}
	o := orm.NewOrm()
	tempTopic, _ := GetTopic(id)
	cid := StrToSlice(tempTopic.Cid)
	topic := &Topic{Id: tid}
	o.QueryTable("topic").Filter("id", tid).One(topic)
	oldAttach := topic.Attachment
	draft := topic.Draft
	_, err = o.Delete(topic)
	if nil != err {
		return
	}
	//删除附件
	if 0 < len(oldAttach) {
		os.Remove(path.Join("attachment", oldAttach))
	}
	//删除评论
	_, err = o.QueryTable("comment").Filter("tid", tid).Delete()
	//分类的文章数减少
	//相应的分类的文章数也要增加
	//删除草稿不用操作类别
	if !draft {
		for _, id := range cid {
			ctgId, err := strconv.ParseInt(id, 10, 64)
			if nil != err {
				return err
			}
			o := orm.NewOrm()
			ctg := &Category{Id: ctgId}
			err = o.Read(ctg)
			if nil != err {
				return err
			} else { //查找到了
				ctg.TopicCount--
				o.Update(ctg)
			}
		}
	}
	return
}

func BatchDeleteTopic(tids []string, pageSize int64) (int64, error) {
	var err error
	for _, id := range tids {
		err = DeleteTopic(id)
		if nil != err {
			return 0, err
		}
	}
	o := orm.NewOrm()
	qt := o.QueryTable("topic").Filter("draft", false)
	totalRecNum, err := qt.Count()
	if nil != err {
		return 0, err
	}
	totalPageNum := (totalRecNum-1)/pageSize + 1
	return totalPageNum, err
}

func Upvote(id string) (int64, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return 0, err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tid}
	err = o.Read(topic)
	if nil == err {
		topic.Upvote++
		o.Update(topic)
	}
	return topic.Upvote, err
}

/*       文章相关操作结束*/

/*       评论相关操作开始*/

func AddComment(id, nickname, email, content, pid string, pageSize int64) (cmt string, err error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return cmt, err
	}
	var cPid int64
	if 0 < len(pid) {
		cPid, err = strconv.ParseInt(pid, 10, 64)
		if nil != err {
			return cmt, err
		}
	}
	comment := &Comment{
		Tid:      tid,
		Nickname: nickname,
		Content:  content,
		Created:  time.Now(),
		Pid:      cPid,
		Email:    email,
	}
	o := orm.NewOrm()
	_, err = o.Insert(comment)
	if nil != err {
		return cmt, err
	}
	//相应的文章的评论数增加
	var count int64
	var totalPageNum int64

	admin := Admin{
		Email:     "@sina.cn",
		EmailHost: "smtp.sina.com",
		EmailPwd:  "",
	}

	if 0 != tid {
		topic := &Topic{Id: tid}
		err = o.Read(topic)
		if nil == err {
			topic.ReplyCount++
			topic.ReplyTime = comment.Created
			topic.ReplyLastUserId = comment.Id
			o.Update(topic)
			//添加回复楼层数
			comment.Floor = topic.ReplyCount
			o.Update(comment)
			//新注册的邮件发送不了
			SendMail(admin, "659986134@qq.com", comment.Nickname+"评论了你的文章<"+topic.Title+">", comment.Content)
		}
	} else { //留言
		qt := o.QueryTable("comment").Filter("tid", 0)
		count, _ = qt.Count()

		qt = qt.Filter("pid", 0)
		totalRecNum, _ := qt.Count()

		totalPageNum = (totalRecNum-1)/pageSize + 1

		if 1 == count {
			comment.Floor = 1
		} else {
			tempCmt := new(Comment)
			err = o.QueryTable("comment").Filter("tid", 0).Exclude("floor", 0).OrderBy("-created").Limit(1).One(tempCmt)
			if nil != err {
				return cmt, err
			}
			//紧跟最后留言的楼层
			comment.Floor = tempCmt.Floor + 1
		}
		o.Update(comment)
		//发送邮件

		SendMail(admin, "659986134@qq.com", comment.Nickname+"给你留言了", comment.Content)
	}

	cmt = "{id:'" + fmt.Sprintf("%d", comment.Id) + "',created:'" + fmt.Sprintf("%v", comment.Created.Format("2006-01-02 15:04:05")) + "',floor:'" + fmt.Sprintf("%d", comment.Floor) + "',count:'" + fmt.Sprintf("%d", count) + "',totalPageNum:'" + fmt.Sprintf("%d", totalPageNum) + "'}"
	return cmt, err
}

func GetAllComment(id string) (comments []*Comment, err error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return nil, err
	}
	comments = make([]*Comment, 0)
	o := orm.NewOrm()
	qt := o.QueryTable("comment")
	_, err = qt.Filter("tid", tid).All(&comments)
	return
}

//暂时只做留言的分页,todo评论的分页
func GetCommentByPage(id string, page *utils.Page) (*utils.Page, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return nil, err
	}
	comments := make([]*Comment, 0)
	o := orm.NewOrm()
	qt := o.QueryTable("comment").Filter("tid", tid)
	page.StartIndex, _ = qt.Count() //总的留言数
	qt = qt.Filter("pid", 0)
	page.TotalRecNum, err = qt.Count() //顶级留言总数
	_, err = qt.Limit(page.PageSize, (page.PageNow-1)*page.PageSize).All(&comments)
	//循环取出回复的留言(非顶级留言)
	comments = GetReply(comments, comments, tid)
	page.PageContent = comments

	page.TotalPageNum = (page.TotalRecNum-1)/page.PageSize + 1
	page.PrevPage = page.PageNow > 1
	page.NextPage = page.PageNow < page.TotalPageNum

	return page, err
}

//循环取出非顶级留言
func GetReply(iterate, store []*Comment, tid int64) []*Comment {
	for _, cmt := range iterate {
		temp := make([]*Comment, 0)
		o := orm.NewOrm()
		_, err := o.QueryTable("comment").Filter("tid", tid).Filter("pid", cmt.Id).All(&temp)
		if nil == err { //有相应的数据
			store = append(store, temp...)
			store = GetReply(temp, store, tid)
		}
	}
	return store
}

func GetNewComment(pageSize int64) (jsonStr string, err error) {
	o := orm.NewOrm()
	comments := make([]*Comment, 0)
	qs := o.QueryTable("comment").Filter("tid__gt", 0).OrderBy("-created")
	_, err = qs.Limit(pageSize).All(&comments, "id", "tid", "content", "nickname")
	if nil != err {
		return
	}
	if 0 < len(comments) {
		topics := make([]*Topic, 0)
		ids := []int64{}
		for _, cmt := range comments {
			ids = append(ids, cmt.Tid)
		}
		o.QueryTable("topic").Filter("id__in", ids).All(&topics, "id", "title")
		topicMap := make(map[int64]string)
		for _, topic := range topics {
			topicMap[topic.Id] = topic.Title
		}
		jsonStr = "{'comments':["
		length := len(comments)
		for i := 0; i < length; i++ {
			tmpCmt := comments[i]
			//格式: xxx发表在<<xxx>> .....
			if i == (length - 1) {
				jsonStr += "{'nickname':'" + tmpCmt.Nickname + "','tid':'" + fmt.Sprintf("%d", tmpCmt.Tid) + "','title':'" + topicMap[tmpCmt.Tid] + "','content':'" + tmpCmt.Content + "'}]}"
			} else {
				jsonStr += "{'nickname':'" + tmpCmt.Nickname + "','tid':'" + fmt.Sprintf("%d", tmpCmt.Tid) + "','title':'" + topicMap[tmpCmt.Tid] + "','content':'" + tmpCmt.Content + "'},"
			}
		}
	}
	return jsonStr, err
}

func DeleteComment(commentId, topicId string, pageSize int64) (jsonStr string, err error) {

	cid, err := strconv.ParseInt(commentId, 10, 64)
	if nil != err {
		return
	}
	tid, err := strconv.ParseInt(topicId, 10, 64)
	if nil != err {
		return
	}
	comment := &Comment{Id: cid}
	o := orm.NewOrm()
	_, err = o.Delete(comment)
	if nil != err {
		return
	}
	//修改相应的文章的评论数/最后评论时间/最后评论用户
	topic := &Topic{Id: tid}

	//删除对此评论的回复 循环删除
	topicComments, err := GetAllComment(topicId)
	allNum := len(topicComments)
	remain := DelReply([]int64{int64(cid)}, topicComments, o)
	delNum := allNum - remain

	err = o.Read(topic)
	if nil == err {
		topic.ReplyCount -= (int64(delNum) + 1)
		if 0 == topic.ReplyCount {
			var t time.Time
			topic.ReplyTime = t
			topic.ReplyLastUserId = 0
		} else {
			//得到最新的评论
			tempComment := &Comment{}
			_, err = o.QueryTable("comment").Filter("tid", tid).OrderBy("-created").Limit(1).All(tempComment)
			if nil != err {
				return
			}
			topic.ReplyTime = tempComment.Created
			topic.ReplyLastUserId = tempComment.Id
		}
		//需判断是不是删除最后一个评论或者直接查找出最新的评论时间(较好实现)
		/*topic.ReplyTime = comment.Created
		topic.ReplyLastUserId = comment.Id*/
		o.Update(topic)
	}
	//评论
	//if 0 == tid {
	qt := o.QueryTable("comment").Filter("tid", tid).Filter("pid", 0)
	totalRecNum, err := qt.Count()
	if nil != err {
		return
	}
	totalPageNum := (totalRecNum-1)/pageSize + 1
	jsonStr = `{"delNum":"` + fmt.Sprintf("%d", (int64(delNum)+1)) + `","totalPageNum":"` + fmt.Sprintf("%d", totalPageNum) + `"}`
	//}
	return
}

func GetTopicByDate(date string, page *utils.Page) (*utils.Page, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qt := o.QueryTable("topic").Filter("draft", false).Filter("created__startswith", date)
	var err error
	page.TotalRecNum, err = qt.Count()
	_, err = qt.Limit(page.PageSize, (page.PageNow-1)*page.PageSize).All(&topics)
	page.TotalPageNum = (page.TotalRecNum-1)/page.PageSize + 1
	page.PrevPage = page.PageNow > 1
	page.NextPage = page.PageNow < page.TotalPageNum
	//转换分类信息 id-->文字
	//先取出所有的分类在比较 一次数据库查询而已
	allCtg, err := GetAllCategory()
	for _, topic := range topics {
		//这里只考虑一个分类，刚开始考虑不周到，如果没有填写分类，应该给个默认值
		cid, _ := strconv.ParseInt(StrToSlice(topic.Cid)[0], 10, 64)
		topic.TempStr = StrToSlice(topic.Cid)[0]
		for _, ctg := range allCtg {
			if cid == ctg.Id {
				topic.Cid = ctg.Title
				break //一篇文章只属于一个分类?
			}
		}
	}
	page.PageContent = topics
	return page, err
}

/*       文章相关操作结束*/

/* $123#$123#*/
func StrToSlice(str string) []string {
	//$123#$123$
	str = strings.TrimSpace(str)
	if 0 < len(str) {
		tempStr := strings.Replace(strings.Replace(str, "#", ",", -1), "$", "", -1)
		//$123#$123$ --> [123 123]
		return strings.Split(strings.TrimSuffix(tempStr, ","), ",") //先把多余的,干掉
	}
	return []string{}
}

//取得所有标签
func GetAllLabel() []string {
	topics := make([]*Topic, 0)
	o := orm.NewOrm()
	o.QueryTable("topic").Filter("draft", false).All(&topics)
	labelMap := make(map[string]int)
	for _, topic := range topics {
		for _, l := range StrToSlice(topic.Label) {
			labelMap[l] = 1
		}
	}
	allLabel := make([]string, 0)
	for label, _ := range labelMap {
		allLabel = append(allLabel, label)
	}
	return allLabel
}

//递归删除评论的回复，相应文章的评论数也要减少
//每次递归都不会共享 topicComments，每个调用栈单独拥有，注意返回值的顺序!!!!!!!!!!!!!!
//range topicComments 并不会每次都取最新的!!!!!!!!!!
//len 会动态计算长度
func DelReply(ids []int64, topicComments []*Comment, tempOrm orm.Ormer) int {
	toDel := []int64{}
	for _, id := range ids {
		for k := 0; k < len(topicComments); {
			cmt := topicComments[k]
			flag := false
			if cmt.Pid == id {
				toDel = append(toDel, cmt.Id)
				tempOrm.QueryTable("comment").Filter("id", cmt.Id).Delete()
				if 0 == k {
					if 1 == len(topicComments) {
						topicComments = make([]*Comment, 0)
					} else {
						topicComments = topicComments[1:]
					}
				} else if (len(topicComments) - 1) == k {
					topicComments = topicComments[:k]
				} else { //删除中间的
					topicComments = append(topicComments[:k], topicComments[k+1:]...)
				}
				flag = true
			}
			//这边搞了好久，注意理清思绪!!!!!!
			if flag {
				k = 0
			} else {
				k++
			}
		}
	}
	if len(toDel) > 0 {
		return DelReply(toDel, topicComments, tempOrm)
	}
	return len(topicComments)
}

//取得博客档案信息
func GetArchive(page *utils.Page) (*utils.Page, error) {
	o := orm.NewOrm()
	allTopic := make([]*Topic, 0)
	archive := make([]string, 0)
	_, err := o.QueryTable("topic").Filter("draft", false).OrderBy("-created").All(&allTopic)
	if nil != err {
		return page, err
	}
	for _, topic := range allTopic {
		createdTime := fmt.Sprintf("%v", topic.Created.Format("2006-01"))
		createdTime = strings.Replace(createdTime, "-", " 年 ", 1) + " 月"
		flag := true
		for _, v := range archive {
			if createdTime == v {
				flag = false
				break
			}
		}
		if flag {
			archive = append(archive, createdTime)
		}
	}

	page.TotalRecNum = int64(len(archive))
	page.TotalPageNum = (page.TotalRecNum-1)/page.PageSize + 1
	page.PrevPage = page.PageNow > 1
	page.NextPage = page.PageNow < page.TotalPageNum
	//每页数量大于总共的数量,全部显示
	if page.PageSize >= page.TotalRecNum {
		page.PageContent = archive[:page.TotalRecNum]
	} else { //可以分页
		//最后一页
		if 0 == page.PageNow-page.TotalPageNum {
			page.PageContent = archive[(page.PageNow-1)*page.PageSize : page.TotalRecNum]
		} else {
			page.PageContent = archive[(page.PageNow-1)*page.PageSize : page.PageSize+(page.PageNow-1)*page.PageSize]
		}
	}
	return page, err
}

/*todo 相关操作 start */

func GetAllTodo() ([]*Todo, error) {
	todos := make([]*Todo, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("todo").All(&todos)
	return todos, err
}

func AddTodo(content string) (string, error) {
	o := orm.NewOrm()
	todo := &Todo{
		Content: content,
		Start:   time.Now(),
	}
	_, err := o.Insert(todo)
	return fmt.Sprintf("%d", todo.Id), err
}

func UpdateStatus(id string) error {
	todoId, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return err
	}
	todo := &Todo{Id: todoId}
	o := orm.NewOrm()
	err = o.Read(todo)
	if nil != err {
		return err
	} else {
		if todo.Done {
			todo.Done = false
		} else {
			todo.Done = true
		}
		o.Update(todo)
	}
	return err
}

/*todo 相关操作 end */

/*简历 相关操作 start */
func AddResume(title, color, content string) (string, error) {
	o := orm.NewOrm()
	resume := &Resume{
		Title:   title,
		Color:   color,
		Content: content,
	}
	_, err := o.Insert(resume)
	return fmt.Sprintf("%d", resume.Id), err
}
func GetResumeById(id string) (*Resume, error) {
	rid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return nil, err
	}
	o := orm.NewOrm()
	qt := o.QueryTable("resume")
	resume := new(Resume)
	err = qt.Filter("id", rid).One(resume)
	return resume, err
}
func GetResume() ([]*Resume, error) {
	o := orm.NewOrm()
	resumes := make([]*Resume, 0) //Not ↓
	_, err := o.QueryTable("resume").Exclude("color__iexact", "WordFile").All(&resumes)
	return resumes, err
}
func GetWordFile() (*Resume, error) {
	o := orm.NewOrm()
	wordFile := new(Resume)
	err := o.QueryTable("resume").Filter("color__iexact", "WordFile").One(wordFile)
	return wordFile, err
}
func EditResume(id, title, color, content string) error {
	rid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return err
	}
	o := orm.NewOrm()
	resume := &Resume{Id: rid}
	err = o.Read(resume)
	if nil != err {
		return err
	} else { //查找到了
		resume.Title = title
		resume.Color = color
		resume.Content = content
		o.Update(resume)
	}
	return nil
}
func DeleteResume(id string) error {
	rid, err := strconv.ParseInt(id, 10, 64)
	if nil != err {
		return err
	}
	resume := &Resume{Id: rid}
	o := orm.NewOrm()
	//删除简历
	err = o.Read(resume)
	if nil != err {
		return err
	} else {
		if "WordFile" == resume.Color {
			os.Remove(path.Join("attachment", resume.Title))
		}
	}
	_, err = o.Delete(resume)
	return err
}

/*简历 相关操作 end */

/*图片 相关操作 start */
func AddAlbum(name, description string) (string, error) {
	o := orm.NewOrm()
	album := &Album{
		Name:        name,
		Description: description,
		Created:     time.Now(),
	}
	_, err := o.Insert(album)
	//todo:创建相册文件夹？
	return `{"id":"` + fmt.Sprintf("%d", album.Id) + `","created":"` + fmt.Sprintf("%v", album.Created.Format("2006-01-02 15:04:05")) + `"}`, err
}

func GetAlbum() ([]*Album, error) {
	o := orm.NewOrm()
	albums := make([]*Album, 0)
	_, err := o.QueryTable("album").All(&albums)
	//设置封面及计算图片数量
	for _, album := range albums {
		photos := make([]*Photo, 0)
		count, err := o.QueryTable("photo").Filter("AlbumId", album.Id).All(&photos)
		if nil != err {
			return nil, err
		}
		album.Num = count
		if 0 != len(photos) {
			album.Cover = "/static/img/photo/" + photos[0].Name
		} else {
			//默认相册封面
			album.Cover = "/static/img/photo/default.jpg"
		}
	}
	return albums, err
}

func AddPhoto(aid, name, description string) error {
	id, err := strconv.ParseInt(aid, 10, 64)
	if nil != err {
		return err
	}
	o := orm.NewOrm()
	photo := &Photo{
		AlbumId:     id,
		Name:        name,
		Description: description,
		Created:     time.Now(),
	}
	_, err = o.Insert(photo)
	return err
}

/*图片 相关操作 end */
func SendMail(admin Admin, to, title, content string) (err error) {
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	header := make(map[string]string)
	//from := mail.Address{"发送人", email}
	from := mail.Address{"博客系统", admin.Email}
	//header["From"] = from.String()
	header["From"] = from.String()
	header["To"] = to
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(title)))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"

	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + b64.EncodeToString([]byte(content))
	//认证
	auth := smtp.PlainAuth("", admin.Email, admin.EmailPwd, admin.EmailHost)
	//发送
	err = smtp.SendMail(admin.EmailHost+":25", auth, admin.Email, []string{to}, []byte(msg))

	return err
}
