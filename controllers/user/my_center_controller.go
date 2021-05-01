package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/auth"
	"ions_zhiliao/utils"
	"strconv"
	"strings"
)

type MyCenterController struct {
	beego.Controller
}

func (m *MyCenterController) Get()  {
	user_id := m.GetSession("id")
	o:= orm.NewOrm()
	user := auth.User{}
	_ = o.QueryTable("sys_user").Filter("id", user_id).One(&user)
	m.Data["user"]= user
	m.TplName = "user/my_center_edit.html"
}

func (m *MyCenterController) Post()  {
	id, _ := m.GetInt("uid")
	username := m.GetString("username")
	old_pwd := m.GetString("old_pwd")
	new_pwd := m.GetString("new_pwd")
	age, _ := m.GetInt("age")
	gender, _ := m.GetInt("gender")
	phone := m.GetString("phone")
	addr := strings.Trim(m.GetString("addr")," ")
	isActive, _ := m.GetInt("is_active")

	phone_int64, _ := strconv.ParseInt(phone,10,64)
	fmt.Println(username,old_pwd,new_pwd,age,gender,phone,addr,isActive)
	o:= orm.NewOrm()
	qs := o.QueryTable("sys_user")
	user_data := auth.User{}
	_ = qs.Filter("id",id).One(&user_data)
	old_pwd_md5 := utils.GetMD5Str(old_pwd)
	ret_map := map[string]interface{}{}
	fmt.Println(old_pwd_md5)
	fmt.Println(user_data.Password)
	if old_pwd_md5 != user_data.Password {
		ret_map["code"] = 10001
		ret_map["msg"] = "原密码错误！"

	}else {
		//user.UserName = username
		//user.Password = utils.GetMD5Str(new_pwd)
		//user.Age = age
		//user.Addr = addr
		//user.Gender = gender
		//user.Phone = phone_int64
		//user.IsActive = isActive
		_, err := qs.Filter("id", id).Update(orm.Params{
			"username":  username,
			"password":  utils.GetMD5Str(new_pwd),
			"age":       age,
			"gender":    gender,
			"phone":     phone_int64,
			"addr":      addr,
			"is_active": isActive,
		})
		if err !=nil {
			logs.Info("修改失败:",err)
			ret_map["code"] = 10001
			ret_map["msg"] = "修改失败！"
		}else {
			ret_map["code"] = 200
			ret_map["msg"] = "修改成功！"
		}
	}
	m.Data["json"] = ret_map
	m.ServeJSON()
}