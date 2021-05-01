package caiwu

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/my_center"
	"ions_zhiliao/utils"
	"math"
	"path"
	"strconv"
	"time"
)

type CaiWuSalarySlipController struct {
	beego.Controller
}

func (c *CaiWuSalarySlipController) Get()  {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_salary_slip")


	var salary_slips []my_center.SalarySlip

	pageSize := 10
	pageNo, err := c.GetInt("page")
	if err != nil {
		pageNo = 1
	}
	var count int64 = 0
	offsetNum := pageSize * (pageNo-1)
	month := c.GetString("month")
	if month != "" {
		count, _ = qs.Filter("pay_date", month).Filter("pay_date", month).Count()
		_, _ = qs.Filter("pay_date",month).Limit(pageSize).Offset(offsetNum).All(&salary_slips)
	}else {
		month = time.Now().Format("2006-01")
		count, _ = qs.Filter("pay_date", month).Count()
		_, _ = qs.Filter("pay_date", month).Limit(pageSize).Offset(offsetNum).All(&salary_slips)
	}
	//_, err := qs.Filter("pay_date", month).All(&salary_slips)
	//if err != nil {
	//	logs.Info("财务中心-工资条错误：",err)
	//}
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

	c.Data["salary_slips"] = salary_slips
	c.Data["count"] = count
	c.Data["pageNo"] = pageNo
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["countPage"] = countPage
	c.Data["page_map"] = page_map
	c.Data["kw"] = month
	c.TplName = "caiwu/salary_slip_list.html"
}

func (c *CaiWuSalarySlipController) ToImportExcel()  {
	c.TplName = "caiwu/salary_slip_import.html"
}

func (c *CaiWuSalarySlipController) DoImportExcel()  {
	f, h, err := c.GetFile("upload_file")
	ret_map := map[string]interface{}{}
	err_data_arr := []string{}
	defer func() {
		f.Close()
	}()
	if err != nil {
		ret_map["code"] = 10001
		ret_map["msg"] = "文件上传失败"
		logs.Info("excel获取失败：",err)
		c.Data["json"] = ret_map
		c.ServeJSON()
		return
	}
	file_name :=  h.Filename
	if path.Ext(file_name) != ".xlsx" {
		ret_map["code"] = 10001
		ret_map["msg"] = "文件类型错误,请选择 *.xlsx 文件 !"
		c.Data["json"] = ret_map
		c.ServeJSON()
		return
	}
	fmt.Println(file_name)
	time_unix_int := time.Now().Unix()
	time_unix_str := strconv.FormatInt(time_unix_int,10)
	fmt.Println(time_unix_str)
	file_path := "upload/salary_slip_upload/"+time_unix_str+"-"+file_name
	err = c.SaveToFile("upload_file", file_path)
	if err != nil {
		ret_map["code"] = 10001
		ret_map["msg"] = "文件保存失败"
		logs.Info("DoImportExcel文件保存失败：",err)
	}else {
		//读取数据并插入数据库
		file, err := excelize.OpenFile(file_path)
		if err != nil {
			ret_map["code"] = 10001
			ret_map["msg"] = "文件获取失败"
			logs.Info("excelize获取文件失败：",err)
			c.Data["json"] = ret_map
			c.ServeJSON()
		}
		rows :=file.GetRows("Sheet1")
        o := orm.NewOrm()
		for i := 1;i<len(rows);i++ {
			card_id := rows[i][2]
			base_pay,_ := strconv.ParseFloat(rows[i][3],64)
			working_day,_  := strconv.ParseFloat(rows[i][4],64)
			days_off,_  := strconv.ParseFloat(rows[i][5],64)
			days_off_no,_  := strconv.ParseFloat(rows[i][6],64)
			reward,_  := strconv.ParseFloat(rows[i][7],64)
			rent_subsidy,_  := strconv.ParseFloat(rows[i][8],64)
			trans_subsidy,_  := strconv.ParseFloat(rows[i][9],64)
			social_security,_  := strconv.ParseFloat(rows[i][10],64)
			house_provident_fund,_  := strconv.ParseFloat(rows[i][11],64)
			personal_income_tax,_  := strconv.ParseFloat(rows[i][12],64)
			fine,_  := strconv.ParseFloat(rows[i][13],64)
			net_salary,_  := strconv.ParseFloat(rows[i][14],64)
			pay_date := rows[i][15]
			salary_slip := my_center.SalarySlip{
				CardId:card_id,
				BasePay:base_pay,
				WorkingDay:working_day,
				DaysOff:days_off,
				DaysOffNo:days_off_no,
				Reward:reward,
				RentSubsidy:rent_subsidy,
				TransSubsidy:trans_subsidy,
				SocialSecurity:social_security,
				HouseProvidentFund:house_provident_fund,
				PersonalIncomeTax:personal_income_tax,
				Fine:fine,
				NetSalary:net_salary,
				PayDate:pay_date,
			}
			//重复导入相同月份；先删除后倒入
			qs := o.QueryTable("sys_salary_slip")
			is_exist := qs.Filter("card_id", card_id).Filter("pay_date", pay_date).Exist()
			if is_exist {
				_, err = qs.Filter("card_id", card_id).Filter("pay_date", pay_date).Delete()
				if err != nil {
					logs.Error("删除相同月份失败：",err)
				}
			}
			_, err = o.Insert(&salary_slip)
			if err != nil {
				//精确到导入失败数据信息
				err_data_arr = append(err_data_arr, card_id)
				logs.Error("当前数据导入失败：",err)

			}
		}
		if len(err_data_arr)<=0 {
			ret_map["code"] = 200
			ret_map["msg"] = "文件上传成功"
		}else {
			ret_map["code"] = 10002
			ret_map["err_data"] = err_data_arr
			ret_map["msg"] = "文件上传成功"
		}
	}
	c.Data["json"] = ret_map
	c.ServeJSON()
}