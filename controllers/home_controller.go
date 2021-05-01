package controllers

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

type HomeController struct {
	beego.Controller
}

func (h *HomeController) Get()  {
	o := orm.NewOrm()
	user_id := h.GetSession("id")
	user_name := h.GetSession("name")
	user := auth.User{Id:user_id.(int)}
	_, _ = o.LoadRelated(&user, "Role")
	auth_arr := []int{}
	for _,role := range user.Role {
		role_data := auth.Role{Id:role.Id}
		_, _ = o.LoadRelated(&role_data, "Auth")
		for _,auth_data := range role_data.Auth {
			auth_arr = append(auth_arr,auth_data.Id)
		}
	}

	qs := o.QueryTable("sys_auth")
	var auths []auth.Auth
	//_, err := qs.Filter("pid", 0).All(&auths)
	_, err := qs.Filter("pid", 0).Filter("id__in",auth_arr).OrderBy("-weight").All(&auths)
	if err != nil {
		logs.Info("获取Auth树节点失败：",err)
	}
	var trees []auth.Tree
	for _,auth_data := range auths {
		pid := auth_data.Id
		tree_data := auth.Tree{Id:auth_data.Id,AuthName:auth_data.AuthName,UrlFor:auth_data.UrlFor,Weight:auth_data.Weight,Children:[]*auth.Tree{}}
		GetChildNode(pid,&tree_data)
		trees = append(trees, tree_data)
	}
	//fmt.Println(trees)
	//for _,tree := range trees{
	//	for _,node := range tree.Children {
	//		fmt.Println(node)
	//	}
	//}
	//消息通知
	qs1 := o.QueryTable("sys_cars_apply")
	var cars_apply []auth.CarsApply
	_, _ = qs1.Filter("user_id", user_id.(int)).Filter("return_status",0).Filter("audit_status",1).Filter("notify_tag",0).RelatedSel().All(&cars_apply)
	cur_time, _ :=time.Parse( "2006-01-02",time.Now().Format("2006-01-02"))
	for _,apply := range cars_apply {
		return_date := apply.ReturnDate
		ret := cur_time.Sub(return_date)
		content := fmt.Sprintf("%s用户，你借的车辆%s,归还时间为%v，已经逾期，请尽快归还",user_name,apply.Cars.Name,return_date.Format("2006-01-02"))
		if ret>0 {
			message_notify := auth.MessageNotify{
				Flag:1,
				Title:"车辆归还逾期",
				Content:content,
				User:&user,
				ReadTag:0,
			}
			_, _ = o.Insert(&message_notify)
		}
		apply.NotifyTag = 1
		_, _ = o.Update(&apply)
	}
	//展示消息
	qs2 := o.QueryTable("sys_message_notify")
	notify_count,_ := qs2.Filter("read_tag",0).Count()
	h.Data["notify_count"] = notify_count
	h.Data["trees"] = trees
	h.Data["user_name"] = user_name
	h.TplName = "index.html"
}

func (h *HomeController) Welcome()  {
	h.TplName = "welcome.html"
}


func GetChildNode(pid int,treenode *auth.Tree)  {
	o :=  orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	var auths []auth.Auth
	_, err := qs.Filter("pid", pid).OrderBy("-weight").All(&auths)
	if err != nil {
		logs.Info("没有子节点")
		return
	}
	fmt.Sprintln("----------",auths)
	for i:=0;i<len(auths);i++{
		pid := auths[i].Id
		tree_data := auth.Tree{Id:auths[i].Id,AuthName: auths[i].AuthName,UrlFor: auths[i].UrlFor,Weight: auths[i].Weight,Children:[]*auth.Tree{}}
		treenode.Children = append(treenode.Children,&tree_data)
		GetChildNode(pid,&tree_data)
	}
	return
}

func (h *HomeController) NotifyList()  {
	o:= orm.NewOrm()
	qs := o.QueryTable("sys_message_notify")
	var nofities []auth.MessageNotify
	pageSize := 10
	pageNo, err := h.GetInt("page")
	if err != nil {
		pageNo = 1
	}
	var count int64 = 0
	offsetNum := pageSize * (pageNo-1)
	kw := h.GetString("kw")
	if kw != "" {
		count, _ = qs.Filter("title__contains", kw).Count()
		_, _ = qs.Filter("title__contains",kw).Limit(pageSize).Offset(offsetNum).RelatedSel().All(&nofities)
	}else {
		count, _ = qs.Count()
		_, _ = qs.Limit(pageSize).Offset(offsetNum).RelatedSel().All(&nofities)
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

	h.Data["count"] = count
	h.Data["pageNo"] = pageNo
	h.Data["prePage"] = prePage
	h.Data["nextPage"] = nextPage
	h.Data["countPage"] = countPage
	h.Data["page_map"] = page_map
	h.Data["kw"] = kw
	h.Data["nofities"] = nofities
	h.TplName = "notify_list.html"
}

func (h *HomeController) ReadNotify()  {
	id, _ := h.GetInt("id")
	o := orm.NewOrm()
	_, err := o.QueryTable("sys_message_notify").Filter("id", id).Update(orm.Params{
		"read_tag": 1,
	})
	if err != nil {
		logs.Error("ReadNotify已读失败！！！")
	}
	h.Redirect(beego.URLFor("HomeController.NotifyList"),302)
}