package news

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/news"
	"ions_zhiliao/utils"
	"math"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get()  {
	 o := orm.NewOrm()
	 var categories []news.Category
	 qs := o.QueryTable("sys_category")

	pageSize := 10
	pageNo, err := c.GetInt("page")
	if err != nil {
		pageNo = 1
	}
	var count int64 = 0
	offsetNum := pageSize * (pageNo-1)
	kw := c.GetString("kw")
	if kw != "" {
		count, _ = qs.Filter("name__contains", kw).Filter("is_delete", 0).Count()
		_, _ = qs.Filter("name__contains",kw).Filter("is_delete", 0).Limit(pageSize).Offset(offsetNum).All(&categories)
	}else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, _ = qs.Filter("is_delete", 0).Limit(pageSize).Offset(offsetNum).All(&categories)
	}

	//总页数
	countPage := int(math.Ceil(float64(count)/float64(pageSize)))
	prePage := 1
	if pageNo == 1 {
		prePage = 1
	}else if pageNo>1 {
		prePage = pageNo - 1
	}
	nextPage := 1
	if pageNo < countPage {
		nextPage = pageNo + 1
	}else if pageNo >= countPage {
		nextPage = pageNo
	}
	page_map := utils.Paginator(pageNo,pageSize,count)

	c.Data["categories"] = categories
	c.Data["count"] = count
	c.Data["pageNo"] = pageNo
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["countPage"] = countPage
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw

	c.TplName = "news/category_list.html"
}

func (c *CategoryController) ToAdd()  {
	c.TplName = "news/category_add.html"
}

func (c *CategoryController) DoAdd()  {
	name := c.GetString("name")
	desc := c.GetString("desc")
	is_active, _ := c.GetInt("is_active")
	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println(is_active)
	o := orm.NewOrm()
	ret_map := map[string]interface{}{}
	is_exist := o.QueryTable("sys_category").Filter("name",name).Exist()
	fmt.Println(is_exist)
	if is_exist {
		ret_map["code"] = 10001
		ret_map["msg"] = "栏目已经存在！"
	}else {
		category_data := news.Category{Name:name,Desc:desc,IsActive:is_active}
		_, err := o.Insert(&category_data)
		if err != nil {
			ret_map["code"] = 10001
			ret_map["msg"] = "栏目添加失败！"
		}else {
			ret_map["code"] = 200
			ret_map["msg"] = "添加成功！"
		}
	}
	c.Data["json"]=ret_map
	c.ServeJSON()
}