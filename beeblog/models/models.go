//使用 beego 的 ORM 创建数据库
package models

import (
	"github.com/Unknwon/com" //通用函数包
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" //数据库驱动（下划线的意思是只执行它的初始化函数，不调用其他的函数）
	"os"                            //注册数据库、驱动、模型之前先判断路径文件和路径是否存在，不存在则创建路径。
	"path"                          //用于获取文件的路径，不包含文件名
	"strconv"
	"time" //Created ，即文章创建的时间需要引用 time 包里的时间函数
)

const (
	_DB_NAME        = "data/beeblog.db" //数据库文件的路径和名称
	_SQLITE3_DRIVER = "sqlite3"         //驱动名称
)

//模型1：创建博客文章分类
type Category struct {
	Id              int64     //当字段名称为 id 并且类型为 int、int32、int64，ORM 会自动认为它是主键，所以不需要专门去声明它为主键
	Title           string    //在不声明长度的情况下，ORM 自动规定它为 255 字节的字符串
	CreatedTime     time.Time `orm:"index"` //文章创建的时间需要引用 time 包里的时间函数。可能会根据文章创建时间查找，所以给它添加索引。
	Views           int64     `orm:"index"` //浏览次数。可能会根据浏览次数排序，所以给它添加索引。
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64     //统计本分类中有多少篇文章
	TopicLastUserId int64     //记录浏览本分类的最后一位用户的 ID。例如：最近一位用户XXX于某年某月某日某时某分某秒对本分类进行操作。
}

//模型2：创建文章
type Topic struct {
	Id               int64
	Uid              int64     //作者 ID，即文章是谁写的。通过 UID 链接到作者的个人信息。
	Title            string    //文章标题
	Category         string    //所属分类
	Content          string    `orm:"size(5000)"` //文章内容。默认是 255 字节，默认值较小。所以指定文章容量的大小。
	Attachment       string    //附件
	CreatedTime      time.Time `orm:"index"` //文章创建时间
	Updated          time.Time `orm:"index"` //文章最后的更新时间
	Views            int64     `orm:"index"` //被访问次数
	Author           string    //作者名称或者昵称
	ReplayTime       time.Time `orm:"index"` //最后回复评论的时间
	ReplayCount      int64     //总的回复数量
	ReplayLastUserId int64     //最后评论的用户的 Id
}

//func AddTopic(title, content string) error {
func AddTopic(title, content, category string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:       title,
		Content:     content,
		CreatedTime: time.Now(),
		Updated:     time.Now(),
	}

	_, err := o.Insert(topic)
	return err
}

//定义一个方法，由主程序来调用它，注册数据库、驱动、模型
//注册数据库
func RegisterDB() {
	//if !Exist(_DB_NAME) {
	if !com.IsExist(_DB_NAME) { //注册数据库之前先判断数据库文件是否存在
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm) //path.Dir(_DB_NAME) 可以取到数据库文件 beeblog.db 所在的路径
		os.Create(_DB_NAME)                          //创建数据库文件，并且不接收任何返回值
	}

	//注册模型
	orm.RegisterModel(new(Category), new(Topic))

	//注册驱动
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite) //格式：驱动名，驱动类型

	//注册默认数据库
	//orm 允许同时操作多个数据库。但是不管有几个数据，强制必须有一个数据库的名称叫做 default 。
	//第3个参数 datasource 这里不使用数据库的真实路径，例如："data/beeblog.db"，而是使用常量 _DB_NAME
	//第4个参数 即数据库最大连接数
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10) //
}

//func AddCategory(name string, flag int) error {
func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name}

	qs := o.QueryTable("category") //进行查询，判断 name 是否已经被使用过
	err := qs.Filter("title", name).One(cate)
	if err == nil { //如果 err 等于 nil，说明 name 已经被使用过
		return err
	}

	_, err = o.Insert(cate)
	if err != nil { //如果错误不为空，说明插入失败
		return err
	}

	return nil
}

//按照 id 删除分类
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func DelTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	_, err = o.Delete(topic)
	return err
}

//获取文章列表。
func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-createdtime").All(&topics) //CreatedTime 在这里全部小写，因为这个字段名称在数据库里都会转化为小写字母
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

//获取所有 Category。返回元素类型是 category 的 slice
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64) //获取到的文章 id 是字符串类型，需要转化为 int64
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func ModifyTopic(tid, title, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}

	return err
}
