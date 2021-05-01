package login

import (
	"fmt"
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/auth"
	"ions_zhiliao/utils"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get()  {
	id,base64,err := utils.GetCaptcha()
	if err != nil {
		ret := fmt.Sprintf("登录get请求，获取验证码错误，错误信息：%v", err)
		logs.Error(ret)
		return
	}
	l.Data["captcha"] = utils.Captcha{Id:id,BS64:base64}
	l.TplName = "login/login.html"
}

func (l *LoginController) Post()  {
	username := l.GetString("username")
	password := l.GetString("password")
	captcha := l.GetString("captcha")
	captcha_id := l.GetString("captcha_id")
	fmt.Println(username,password,captcha,captcha_id)
	is_ok := utils.VerityCaptcha(captcha_id,captcha)
	md5_pwd := utils.GetMD5Str(password)
	userinfo := auth.User{}
	ret_map := map[string]interface{}{}
	if is_ok {
		o := orm.NewOrm()
		isExist := o.QueryTable("sys_user").Filter("user_name", username).Filter("password", md5_pwd).Exist()
		fmt.Println(isExist)
		if isExist {
			error := o.QueryTable("sys_user").Filter("user_name", username).Filter("password", md5_pwd).One(&userinfo)
			if error != nil {
				fmt.Println(error)
			}
			if userinfo.IsActive == 0 {
				logs.Info("该用户已停用，请联系管理员")
				ret_map["code"] = 10001
				ret_map["msg"] = "该用户已停用，请联系管理员"
				l.Data["json"] = ret_map
			}else {
				l.SetSession("id", userinfo.Id)
				l.SetSession("name", userinfo.UserName)
				ret_map["code"] = 200
				ret_map["msg"] = "登录成功"
				l.Data["json"] = ret_map
			}
		}else {
			ret := fmt.Sprintf("登录post请求，用户名密码错误，错误信息：username：%s,password:%s", username,password)
			logs.Info(ret)
			ret_map["code"] = 10001
			ret_map["msg"] = "账号或密码错误"
			l.Data["json"] = ret_map
		}
	}else {
		ret_map["code"] = 10001
		ret_map["msg"] = "验证码错误"
		l.Data["json"] = ret_map
	}
	fmt.Println(is_ok)
	l.ServeJSON()
}

func (l *LoginController) ChangeCaptcha()  {
	message := map[string]interface{}{}
	id,base64,err := utils.GetCaptcha()
	if err != nil {
		ret := fmt.Sprintf("生成验证码失败，错误信息：id：%s,error:%s",id, err)
		logs.Info(ret)
		message["msg"] = "生成失败"
		message["Code"] = 404
		l.Data["json"] = message
	}else {
		l.Data["json"] = utils.Captcha{Id:id,BS64:base64,Code:200}
	}
	fmt.Println("=======更新验证码========")
	l.ServeJSON()
}

func (l *LoginController) LogOut()  {
	l.DestroySession()
	//l.DelSession("id")
	//l.DelSession("name")
	l.Redirect(beego.URLFor("LoginController.Get"),302)
}

