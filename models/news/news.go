package news

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	Id int `orm:"pk;auto"`
	Name string `orm:"size(64);description(分类名称)"`
	Desc string `orm:"size(255);description(描述)"`
	IsActive int `orm:"description(1启用，0停用);default(1)"`
	IsDelete int `orm:"description(1删除，0未删除);default(0)"`
	CreateTime time.Time `orm:"type(datetime);auto_now;description(创建时间)"`
	News []*News `orm:"reverse(many)"`
}

type News struct {
	Id int `orm:"pk;auto"`
	Title string `orm:"size(64);description(新闻标题);type(text)"`
	Content string `orm:"size(64);description(新闻内容)"`
	Category *Category `orm:"rel(fk)"`
	IsActive int `orm:"description(1启用，0停用);default(1)"`
	IsDelete int `orm:"description(1删除，0未删除);default(0)"`
	CreateTime time.Time `orm:"type(datetime);auto_now;description(创建时间)"`
}

func (c *Category) TableName() string {
	return "sys_category"
}

func (n *News) TableName() string {
	return "sys_news"
}

func init()  {
	orm.RegisterModel(new(Category),new(News))
}