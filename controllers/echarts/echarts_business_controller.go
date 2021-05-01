package echarts

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type EchartsBusinessController struct {
	beego.Controller
}

func (e *EchartsBusinessController) Get()  {
	e.TplName = "echarts/echarts_business.html"
}

func (e *EchartsBusinessController) GetBusinessChart()  {
	var caiwu_date orm.ParamsList
	var student_increase orm.ParamsList
	o := orm.NewOrm()
	_, _ = o.Raw("select cai_wu_date from sys_caiwu_data").ValuesFlat(&caiwu_date)
	_, _ = o.Raw("select student_increase from sys_caiwu_data").ValuesFlat(&student_increase)

	map_data := map[string]interface{}{}

	map_data["caiwu_date"]  = caiwu_date
	map_data["student_increase"] = student_increase

	e.Data["json"] = map_data
	e.ServeJSON()
}