package user

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/auth"
	"ions_zhiliao/models/my_center"
	"time"
)

type SalarySlipController struct {
	beego.Controller
}

func (s *SalarySlipController) Get()  {
	month := s.GetString("month")
	if month == ""{
		month = time.Now().Format("2006-01")
	}
	id := s.GetSession("id")
	o:= orm.NewOrm()
	user := auth.User{}
	salary_slip := my_center.SalarySlip{}
	err := o.QueryTable("sys_user").Filter("id", id).One(&user)
	if err != nil {
		logs.Info("工资条获取个人信息失败",err)
	}else {
		card_id := user.CardId
		_ = o.QueryTable("sys_salary_slip").Filter("card_id", card_id).Filter("pay_date", month).One(&salary_slip)
	}
	s.Data["salary_slip"]=salary_slip
	s.TplName = "user/salary_slip_list.html"
}

func (s *SalarySlipController) Detail()  {
	id := s.GetString("id")
	o:= orm.NewOrm()
	salary_slip := my_center.SalarySlip{}
	_ = o.QueryTable("sys_salary_slip").Filter("id", id).One(&salary_slip)
	s.Data["salary_slip"]=salary_slip
	s.TplName = "user/salary_slip_detail.html"
}