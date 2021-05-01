package caiwu

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/caiwu"
	"ions_zhiliao/utils"
	"math"
	"path"
	"strconv"
	"time"
)

type CaiWuEchartDataController struct {
	beego.Controller
}

func (c *CaiWuEchartDataController) Get()  {
	o:=orm.NewOrm()
	var caiwu_datas []caiwu.CaiwuData
	qs := o.QueryTable("sys_caiwu_data")

	pageSize := 10
	pageNo, err := c.GetInt("page")
	if err != nil {
		pageNo = 1
	}
	var count int64 = 0
	offsetNum := pageSize * (pageNo-1)
	month := c.GetString("month")
	if month != "" {
		count, _ = qs.Filter("cai_wu_date", month).Count()
		_, _ = qs.Filter("cai_wu_date",month).Limit(pageSize).Offset(offsetNum).All(&caiwu_datas)
	}else {
		month = time.Now().Format("2006-01")
		count, _ = qs.Filter("cai_wu_date", month).Count()
		_, _ = qs.Filter("cai_wu_date", month).Limit(pageSize).Offset(offsetNum).All(&caiwu_datas)
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

	c.Data["caiwu_datas"] = caiwu_datas
	c.Data["count"] = count
	c.Data["pageNo"] = pageNo
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["countPage"] = countPage
	c.Data["page_map"] = page_map
	c.Data["kw"] = month

	c.TplName = "caiwu/echart_data_list.html"
}

func (c *CaiWuEchartDataController) ToImportExcel()  {
	c.TplName = "caiwu/echart_data_import.html"
}

func (c *CaiWuEchartDataController) DoImportExcel()  {
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
	file_path := "upload/echart_data_upload/"+time_unix_str+"-"+file_name
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
			caiwu_date := rows[i][0]
			sales_volume,_ := strconv.ParseFloat(rows[i][1],64)
			student_increase := utils.StrToInt(rows[i][2])
			django := utils.StrToInt(rows[i][3])
			vue_django := utils.StrToInt(rows[i][4])
			celery := utils.StrToInt(rows[i][5])
			echart_data := caiwu.CaiwuData{
				CaiWuDate:caiwu_date,
				SalesVolume:sales_volume,
				StudentIncrease:student_increase,
				Django:django,
				VueDjango:vue_django,
				Celery:celery,
			}
			//重复导入相同月份；先删除后倒入
			qs := o.QueryTable("sys_caiwu_data")
			is_exist := qs.Filter("cai_wu_date", caiwu_date).Exist()
			if is_exist {
				_, err = qs.Filter("cai_wu_date", caiwu_date).Delete()
				if err != nil {
					logs.Error("删除相同月份失败：",err)
				}
			}
			_, err = o.Insert(&echart_data)
			if err != nil {
				//精确到导入失败数据信息
				err_data_arr = append(err_data_arr, caiwu_date)
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