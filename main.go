package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "ions_zhiliao/models/auth"
	_ "ions_zhiliao/models/my_center"
	_ "ions_zhiliao/models/caiwu"
	_ "ions_zhiliao/models/news"
	_ "ions_zhiliao/routers"
	"ions_zhiliao/utils"
)

func init()  {
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	database := beego.AppConfig.String("database")
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	conn := username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&loc=Local"
	_ = orm.RegisterDataBase("default", "mysql", conn)
	//orm.Debug = true
}

func main() {
	orm.RunCommand()
	beego.BConfig.WebConfig.Session.SessionOn = true

	//登录拦截
	beego.InsertFilter("/admin/*",beego.BeforeRouter,utils.LoginFilter)
	// 日志
	_ = logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/ions.log","separate":["error","info"]}`)
	logs.SetLogFuncCallDepth(3)
	beego.SetStaticPath("/upload","upload")
	// 打印sql
	orm.Debug = true
	//设置静态文件
	beego.SetStaticPath("/upload","upload")
	beego.Run()
}

