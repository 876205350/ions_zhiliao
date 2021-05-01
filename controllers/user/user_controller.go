package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/auth"
	"ions_zhiliao/utils"
	"math"
	"strconv"
	"strings"
)

type UserController struct {
	beego.Controller
}
/*
用户列表
 */
func (u *UserController) List()  {

	//o := orm.NewOrm()
	//
	//qs := o.QueryTable("sys_user")
	//var users []user.User
	//pageSize := 2
	//pageNo, err := u.GetInt("page")
	//if err != nil {
	//	pageNo = 1
	//}
	//offsetNum := pageSize * (pageNo-1)
	//count, _ := qs.Filter("is_delete",0).Count()
	////总页数
	//countPage := int(math.Ceil(float64(count)/float64(pageSize)))
	//_, _ = qs.Filter("is_delete",0).Limit(pageSize).Offset(offsetNum).All(&users)
	///*
	//分页
	//	当前第几页  offset  limit
	//		1		  0 	 2    2*(1-1)
	//		2		  2		 2    2*(2-1)
	// */
	//prePage := 1
	//if pageNo == 1 {
	//	prePage = 1
	//}else if pageNo>1 {
	//	prePage = pageNo - 1
	//}
	//nextPage := 1
	//if pageNo < countPage {
	//	nextPage = pageNo + 1
	//}else if pageNo >= countPage {
	//	nextPage = pageNo
	//}
	o := orm.NewOrm()

	qs := o.QueryTable("sys_user")
	var users []auth.User
	pageSize := 10
	pageNo, err := u.GetInt("page")
	if err != nil {
		pageNo = 1
	}
	var count int64 = 0
	offsetNum := pageSize * (pageNo-1)
	kw := u.GetString("kw")
	if kw != "" {
		count, _ = qs.Filter("is_delete",0).Filter("user_name__contains",kw).Count()
		_, _ = qs.Filter("is_delete",0).Filter("user_name__contains",kw).Limit(pageSize).Offset(offsetNum).All(&users)
	}else {
		count, _ = qs.Filter("is_delete",0).Count()
		_, _ = qs.Filter("is_delete",0).Limit(pageSize).Offset(offsetNum).All(&users)
	}
	//总页数
	countPage := int(math.Ceil(float64(count)/float64(pageSize)))

	/*
		分页
			当前第几页  offset  limit
				1		  0 	 2    2*(1-1)
				2		  2		 2    2*(2-1)
	*/
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
	u.Data["users"] = users
	u.Data["count"] = count
	u.Data["pageNo"] = pageNo
	u.Data["prePage"] = prePage
	u.Data["nextPage"] = nextPage
	u.Data["countPage"] = countPage
	u.Data["page_map"] = page_map
	u.Data["kw"] = kw
	u.TplName = "user/user_list.html"
}
/*
添加用户页
 */
func (u *UserController) ToAdd()  {
	u.TplName = "user/user_add.html"
}
/*
添加用户
*/
func (u *UserController) DoAdd()  {
	username := u.GetString("username")
	password := u.GetString("password")
	age, _ := u.GetInt("age")
	gender, _ := u.GetInt("gender")
	phone := u.GetString("phone")
	addr := u.GetString("addr")
	isActive, _ := u.GetInt("isActive")

	new_password := utils.GetMD5Str(password)
	phone_int64, _ := strconv.ParseInt(phone,10,64)
	fmt.Println(username,age,password,gender,phone,addr,isActive)
	o := orm.NewOrm()
	user_data := auth.User{UserName: username,Password:new_password,Age:age,Gender:gender,Phone:phone_int64,Addr:addr,IsActive:isActive}
	_, err := o.Insert(&user_data)
	ret_map := map[string]interface{}{}
	if err != nil {
		fmt.Println(err)
		ret_map["code"]=10001
		ret_map["msg"]="添加失败,重新操作！"
		u.Data["json"] = ret_map
	}else {
		ret_map["code"]=200
		ret_map["msg"]="添加成功！"
		u.Data["json"] = ret_map
	}
	u.ServeJSON()
}
/*
启用，停用
*/
func (u *UserController) IsActive()  {
	is_active, _ := u.GetInt("is_active")
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	logs.Info("状态变更",is_active,id)
	qs := o.QueryTable("sys_user").Filter("id",id)
	ret_map := map[string]interface{}{}
	if is_active == 1 {
		_, err := qs.Update(orm.Params{
			"is_active":0,
		})
		if err != nil {
			ret_map["code"]=10001
			ret_map["msg"]="操作失败,重新操作！"
			u.Data["json"] = ret_map
		}else {
			ret_map["code"]=200
			ret_map["msg"]="停用成功！"
			u.Data["json"] = ret_map
		}
	}else if is_active == 0{
		_, err := qs.Update(orm.Params{
			"is_active":1,
		})
		if err != nil {
			ret_map["code"]=10001
			ret_map["msg"]="操作失败,重新操作！"
			u.Data["json"] = ret_map
		}else {
			ret_map["code"]=200
			ret_map["msg"]="启用成功！"
			u.Data["json"] = ret_map
		}
	}
	fmt.Println(is_active,id)
	u.ServeJSON()
}
/*
删除用户
*/
func (u *UserController) Delete()  {
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	_, _ = o.QueryTable("sys_user").Filter("id", id).Update(orm.Params{
		"is_delete":1,
	})
	u.Redirect(beego.URLFor("UserController.List"),302)
}

func (u *UserController) ResetPassword()  {
	id, _ := u.GetInt("id")
	password := utils.GetMD5Str("123456789")
	o := orm.NewOrm()
	_, _ = o.QueryTable("sys_user").Filter("id", id).Update(orm.Params{
		"password":password,
	})
	u.Redirect(beego.URLFor("UserController.List"),302)
}

func (u *UserController) ToUpdate()  {
	id, _ := u.GetInt("id")
	o := orm.NewOrm()
	user_date := auth.User{}
	_ = o.QueryTable("sys_user").Filter("id", id).One(&user_date)
	u.Data["user"]=user_date
	u.TplName = "user/user_edit.html"
}

func (u *UserController) DoUpdate()  {
	id, _ := u.GetInt("id")
	username := u.GetString("username")
	age, _ := u.GetInt("age")
	gender, _ := u.GetInt("gender")
	phone := u.GetString("phone")
	addr := u.GetString("addr")
	isActive, _ := u.GetInt("isActive")

	phone_int64, _ := strconv.ParseInt(phone,10,64)
	fmt.Println(username,age,gender,phone,addr,isActive)
	o := orm.NewOrm()
	ret_map := map[string]interface{}{}
	_, err := o.QueryTable("sys_user").Filter("id", id).Update(orm.Params{
		"user_name": username,
		"age":age,
		"gender":gender,
		"phone":phone_int64,
		"addr":addr,
		"is_active":isActive,
	})
	if err!=nil {
		ret_map["code"] = 10001
		ret_map["msg"]="更新失败！"
		u.Data["json"] = ret_map
	}else {
		ret_map["code"]=200
		ret_map["msg"]="更新成功！"
		u.Data["json"] = ret_map
	}
	u.ServeJSON()
}

func (u *UserController) MuliDelete()  {
	ids := u.GetString("ids")
	fmt.Println("798787987",ids)
	id_arr := strings.Split(ids,",")
	fmt.Println("%T\n",id_arr)
	o := orm.NewOrm()
	ret_map := map[string]interface{}{}
	qs := o.QueryTable("sys_user")
	//思路一
	//for _,v := range id_arr {
	//	fmt.Println(v)
	//	id_int := utils.StrToInt(v)
	//	_, _ = qs.Filter("id", id_int).Update(orm.Params{
	//		"is_delete": 1,
	//	})
	//}
	//思路二
	_, err := qs.Filter("id__in", id_arr).Update(orm.Params{
		"is_delete": 1,
	})
	if err != nil {
		ret_map["code"]=10001
		ret_map["msg"]="批量删除失败！"
		u.Data["json"] = ret_map
	}else {
		ret_map["code"] = 200
		ret_map["msg"] = "批量删除成功！"
		u.Data["json"] = ret_map
	}
	u.ServeJSON()
}