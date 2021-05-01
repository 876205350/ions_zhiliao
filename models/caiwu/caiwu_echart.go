package caiwu

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type CaiwuData struct {
	Id int `orm:"pk;auto"`
	CaiWuDate string `orm:"size(32);description(财务月份)"`
	SalesVolume float64 `orm:"description(本月销售额);digits(10);decimals(2)"`
	StudentIncrease int `orm:"description(学员增加数)"`
	Django int `orm:"description(django课程卖出数量)"`
	VueDjango int `orm:"description(vue+django课程卖出数量)"`
	Celery int `orm:"description(celery课程卖出数量)"`
	CreateTime time.Time `orm:"type(datetime);auto_now;description(创建时间)"`
}

func (c *CaiwuData) TableName() string {
	return "sys_caiwu_data"
}

func init()  {
	orm.RegisterModel(new(CaiwuData))
}