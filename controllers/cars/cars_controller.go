package cars

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/auth"
	"ions_zhiliao/utils"
	"math"
)

type CarsController struct {
	beego.Controller
}

func (c  *CarsController) Get(){
	o := orm.NewOrm()
	var cars_data []auth.Cars
	qs := o.QueryTable("sys_cars")
	//qs.Filter("is_delete",0).RelatedSel().All(&cars_data)
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
		_, _ = qs.Filter("name__contains",kw).Filter("is_delete", 0).Limit(pageSize).Offset(offsetNum).RelatedSel().All(&cars_data)
	}else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, _ = qs.Filter("is_delete", 0).Limit(pageSize).Offset(offsetNum).RelatedSel().All(&cars_data)
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

	c.Data["count"] = count
	c.Data["pageNo"] = pageNo
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["countPage"] = countPage
	c.Data["page_map"] = page_map
	c.Data["kw"] = kw
	c.Data["cars_data"] = cars_data
	c.TplName = "cars/cars_list.html"
}

func (c *CarsController) ToAdd() {
	o := orm.NewOrm()
	var cars_brand []auth.CarBrand
	_, _ = o.QueryTable("sys_cars_brand").Filter("is_delete", 0).All(&cars_brand)
	c.Data["cars_brand"] = cars_brand
	c.TplName = "cars/cars_add.html"
}

func (c *CarsController) DoAdd() {
	name := c.GetString("name")
	cars_brand_id,_ := c.GetInt("cars_brand_id")
	is_active, _ := c.GetInt("is_active")
	fmt.Println(name)
	fmt.Println(cars_brand_id)
	fmt.Println(is_active)
	o := orm.NewOrm()
	cars_brand := auth.CarBrand{Id:cars_brand_id}
	ret_map := make(map[string]interface{},5)
	is_exist := o.QueryTable("sys_cars").Filter("name",name).Exist()
	if is_exist {
		ret_map["code"] = 10001
		ret_map["msg"] = "车辆名称重复！"
	}else {
		cars_data := auth.Cars{
			Name:name,
			CarBrand:&cars_brand,
			Status:1,//由于数据库默认值无效
			IsActive:is_active,
		}
		_, err := o.Insert(&cars_data)

		if err != nil {
			ret_map["code"] = 10001
			ret_map["msg"] = "添加失败！"
		}else {
			ret_map["code"] = 200
			ret_map["msg"] = "添加成功！"
		}
	}
	c.Data["json"] = ret_map
	c.ServeJSON()
}
