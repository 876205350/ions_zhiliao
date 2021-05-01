package my_center

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SalarySlip struct {
	Id int `orm:"pk;auto"`
	CardId string `orm:"column(card_id);size(64);description(员工工号)"`
	BasePay float64 `orm:"description(基本工资);digits(12);decimals(2)"`
	WorkingDay float64 `orm:"description(工作天数);digits(4);decimals(2)"`
	DaysOff float64 `orm:"description(请假天数);digits(4);decimals(2)"`
	DaysOffNo float64 `orm:"description(调休天数);digits(4);decimals(2)"`
	Reward float64 `orm:"description(奖金);digits(8);decimals(2)"`
	RentSubsidy float64 `orm:"description(住房补贴);digits(6);decimals(2)"`
	TransSubsidy float64 `orm:"description(交通补贴);digits(6);decimals(2)"`
	SocialSecurity float64 `orm:"description(社保);digits(6);decimals(2)"`
	HouseProvidentFund float64 `orm:"description(住房公积金);digits(6);decimals(2)"`
	PersonalIncomeTax float64 `orm:"description(个税);digits(6);decimals(2)"`
	Fine float64 `orm:"description(罚金);digits(8);decimals(2)"`
	NetSalary float64 `orm:"description(实发工资);digits(12);decimals(2)"`
	PayDate string `orm:"column(pay_date);size(64);description(工资月份)"`
	CreateTime time.Time `orm:"type(datetime);auto_now;description(创建时间)"`
}

func (s *SalarySlip) TableName() string {
	return "sys_salary_slip"
}

func init()  {
	orm.RegisterModel(new(SalarySlip))
}
