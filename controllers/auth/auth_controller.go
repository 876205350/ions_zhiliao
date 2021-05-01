package auth

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

type AuthController struct {
	beego.Controller
}

func (a *AuthController) List() {
	//每页显示的条数
	o := orm.NewOrm()

	qs := o.QueryTable("sys_auth")
	var auths []auth.Auth
	pageSize := 10
	pageNo, err := a.GetInt("page")
	if err != nil {
		pageNo = 1
	}
	var count int64 = 0
	offsetNum := pageSize * (pageNo-1)
	kw := a.GetString("kw")
	if kw != "" {
		count, _ = qs.Filter("is_delete", 0).Filter("auth_name__contains", kw).Count()
		_, _ = qs.Filter("is_delete",0).Filter("auth_name__contains",kw).Limit(pageSize).Offset(offsetNum).All(&auths)
	}else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, _ = qs.Filter("is_delete",0).Limit(pageSize).Offset(offsetNum).All(&auths)
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
	a.Data["auths"] = auths
	a.Data["count"] = count
	a.Data["pageNo"] = pageNo
	a.Data["prePage"] = prePage
	a.Data["nextPage"] = nextPage
	a.Data["countPage"] = countPage
	a.Data["page_map"] = page_map
	a.Data["kw"] = kw
	a.TplName = "auth/auth-list.html"
}

func (a *AuthController) ToAdd()  {
	var auths []auth.Auth
	o:=orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	_, err := qs.Filter("is_delete", 0).All(&auths)
	if err != nil {
		logs.Info("获取权限数据错误:",err)
	}
	a.Data["auths"] = auths
	a.TplName = "auth/auth-add.html"
}

func (a *AuthController) DoAdd()  {
	pid, err := a.GetInt("pid")
	if err != nil {
		logs.Info("父级id获取失败：",err)
	}
	auth_name := a.GetString("auth_name")
	url_for := a.GetString("url_for")
	desc := a.GetString("desc")
	is_active, err := a.GetInt("is_active")
	if err != nil {
		logs.Info("状态获取失败：",err)
	}
	weight, err := a.GetInt("weight")
	if err != nil {
		logs.Info("权重获取失败：",err)
	}
	fmt.Println("==========",pid,auth_name,url_for,desc,is_active,weight)
	o := orm.NewOrm()
	auth_data := auth.Auth{AuthName:auth_name,UrlFor:url_for,Pid:pid,Desc:desc,CreateTime:time.Now(),IsActive:is_active,Weight:weight}
	ret_map := map[string]interface{}{}
	id, err := o.Insert(&auth_data)
	if err != nil {
		fmt.Println(err)
		ret_map["code"]=10001
		ret_map["msg"]="添加失败,重新操作！"
		a.Data["json"] = ret_map
	}else {
		fmt.Println(id)
		ret_map["code"]=200
		ret_map["msg"]="添加成功！"
		a.Data["json"] = ret_map
	}
	a.ServeJSON()
}


