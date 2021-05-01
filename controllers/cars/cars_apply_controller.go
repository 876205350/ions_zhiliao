package cars

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/auth"
	"ions_zhiliao/utils"
	"math"
	"time"
)

type CarsApplyController struct {
	beego.Controller
}

func (c *CarsApplyController) Get() {
	o := orm.NewOrm()
	var cars_data []auth.Cars
	qs := o.QueryTable("sys_cars")
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
	c.TplName = "cars/cars_apply_list.html"
}

func (c *CarsApplyController) ToApply() {
	id, _ := c.GetInt("id")
	c.Data["id"] = id
	c.TplName = "cars/cars_apply.html"
}

func (c *CarsApplyController) DoApply() {
	reason := c.GetString("reason")
	destination := c.GetString("destination")
	return_date := c.GetString("return_date")
	return_date_new,_ := time.Parse("2006-01-02",return_date)
	cars_id, _ := c.GetInt("cars_id")
	uid := c.GetSession("id")
	//interface -> int  断言函数.(int)
	user := auth.User{Id:uid.(int)}
	cars_data := auth.Cars{Id: cars_id}
	fmt.Println(reason)
	fmt.Println(destination)
	fmt.Println(return_date)
	fmt.Println(cars_id)
	o := orm.NewOrm()
	//默认值ReturnStatus:0,AuditStatus:3,IsActive:1,
	cars_apply := auth.CarsApply{
		User:&user,
		Cars:&cars_data,
		Reason:reason,
		Destination:destination,
		ReturnDate:return_date_new,
		ReturnStatus:0,
		AuditStatus:3,
		IsActive:1,
	}
	ret_map := make(map[string]interface{},5)
	_, err := o.Insert(&cars_apply)
	if err != nil {
		ret_map["code"] = 10001
		ret_map["msg"] = "申请提交失败！"
	}else {
		_, err = o.QueryTable("sys_cars").Filter("id", cars_id).Update(orm.Params{
			"status": 2,
		})
		if err != nil {
			logs.Info("CarsApplyController车辆状态更新失败！")
		}
		ret_map["code"] = 200
		ret_map["msg"] = "申请提交成功！"
	}
	c.Data["json"] = ret_map
	c.ServeJSON()
}

func (c *CarsApplyController) MyApply() {
	o := orm.NewOrm()
	uid := c.GetSession("id")
	qs := o.QueryTable("sys_cars_apply")
	var cars_apply []auth.CarsApply
	pageSize := 10
	pageNo, err := c.GetInt("page")
	if err != nil {
		pageNo = 1
	}
	var count int64 = 0
	offsetNum := pageSize * (pageNo-1)
	kw := c.GetString("kw")
	if kw != "" {
		count, _ = qs.Filter("Cars__name__contains", kw).Filter("is_delete", 0).Filter("user_id",uid.(int)).Count()
		_, _ = qs.Filter("Cars__name__contains",kw).Filter("is_delete", 0).Filter("user_id",uid.(int)).Limit(pageSize).Offset(offsetNum).RelatedSel().All(&cars_apply)
	}else {
		count, _ = qs.Filter("is_delete", 0).Filter("user_id",uid.(int)).Count()
		_, _ = qs.Filter("is_delete", 0).Filter("user_id",uid.(int)).Limit(pageSize).Offset(offsetNum).RelatedSel().All(&cars_apply)
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
	c.Data["cars_apply"] = cars_apply
	c.TplName = "cars/my_apply_list.html"
}

func (c *CarsApplyController) AuditApply()  {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars_apply")
	var cars_apply []auth.CarsApply
	pageSize := 10
	pageNo, err := c.GetInt("page")
	if err != nil {
		pageNo = 1
	}
	var count int64 = 0
	offsetNum := pageSize * (pageNo-1)
	kw := c.GetString("kw")
	if kw != "" {
		count, _ = qs.Filter("Cars__name__contains", kw).Filter("is_delete", 0).Count()
		_, _ = qs.Filter("Cars__name__contains",kw).Filter("is_delete", 0).Limit(pageSize).Offset(offsetNum).RelatedSel().All(&cars_apply)
	}else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, _ = qs.Filter("is_delete", 0).Limit(pageSize).Offset(offsetNum).RelatedSel().All(&cars_apply)
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
	c.Data["cars_apply"] = cars_apply
	c.TplName = "cars/audit_apply_list.html"
}

func (c *CarsApplyController) ToAuditApply() {
	id,_ := c.GetInt("id")
	cars_apply := auth.CarsApply{}
	o := orm.NewOrm()
	_ = o.QueryTable("sys_cars_apply").Filter("id", id).RelatedSel().One(&cars_apply)
	c.Data["cars_apply"] = cars_apply
	c.TplName = "cars/audit_apply.html"
}

func (c *CarsApplyController) DoAuditApply()  {
	option := c.GetString("option")
	id, _ := c.GetInt("id")
	cars_id, _ := c.GetInt("cars_id")
	audit_status, _ := c.GetInt("audit_status")
	fmt.Println(option)
	fmt.Println(id)
	fmt.Println(audit_status)
	o := orm.NewOrm()
	_, _ = o.QueryTable("sys_cars_apply").Filter("id", id).Update(orm.Params{
		"audit_option": option,
		"audit_status":audit_status,
	})
	ret_map := make(map[string]interface{},5)
	if audit_status == 2 {
		_, _ = o.QueryTable("sys_cars").Filter("id", cars_id).Update(orm.Params{
			"status": 1,
		})
		ret_map["code"] = 10001
		ret_map["msg"] = "审核未通过！"
	}else {
		ret_map["code"] = 200
		ret_map["msg"] = "审核通过！"
	}
	c.Data["json"] = ret_map
	c.ServeJSON()
}

func (c *CarsApplyController) DoReturn()  {
	id, _ := c.GetInt("id")
	fmt.Println(id)
	o := orm.NewOrm()
	qs := o.QueryTable("sys_cars_apply")
	_, err := qs.Filter("id", id).Update(orm.Params{
		"return_status": 1,
	})

	if err != nil {
		logs.Error("DoReturn归还失败！")
	}else {
		var cars_apply auth.CarsApply
		_ = qs.Filter("id", id).One(&cars_apply)
		_, err = o.QueryTable("sys_cars").Filter("id", cars_apply.Cars.Id).Update(orm.Params{
			"status": 1,
		})
		if err != nil {
			logs.Error("DoReturn归还失败！")
		}else {
			logs.Info("归还成功！")
		}
	}
	c.Redirect(beego.URLFor("CarsApplyController.MyApply"),302)
}