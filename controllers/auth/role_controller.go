package auth

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/auth"
	"ions_zhiliao/utils"
	"math"
	"strings"
	"time"
)

type RoleController struct {
	beego.Controller
}

func (r *RoleController) List()  {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_role")
	var roles []auth.Role
	pageSize := 10
	pageNo, err := r.GetInt("page")
	if err != nil {
		pageNo = 1
	}
	var count int64 = 0
	offsetNum := pageSize * (pageNo-1)
	kw := r.GetString("kw")
	if kw != "" {
		count, _ = qs.Filter("is_delete", 0).Filter("role_name__contains", kw).Count()
		_, err = qs.Filter("is_delete",0).Filter("role_name__contains",kw).Limit(pageSize).Offset(offsetNum).All(&roles)
		if err != nil {
			logs.Info("角色查询失败：",err)
		}
	}else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, err = qs.Filter("is_delete",0).Limit(pageSize).Offset(offsetNum).All(&roles)
		if err != nil {
			logs.Info("角色查询失败：",err)
		}
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
	r.Data["roles"] = roles
	r.Data["count"] = count
	r.Data["pageNo"] = pageNo
	r.Data["prePage"] = prePage
	r.Data["nextPage"] = nextPage
	r.Data["countPage"] = countPage
	r.Data["page_map"] = page_map
	r.Data["kw"] = kw
	r.TplName = "auth/role-list.html"
}

func (r *RoleController) ToAdd()  {
	r.TplName = "auth/role-add.html"
}

func (r *RoleController) DoAdd()  {
	role_name := r.GetString("role_name")
	desc := r.GetString("desc")
	is_active, err := r.GetInt("is_active")
	if err != nil {
		logs.Info("状态获取失败：",err)
	}
	fmt.Println(role_name,desc,is_active)
	o := orm.NewOrm()
	role_data := auth.Role{RoleName:role_name,Desc:desc,CreateTime:time.Now(),IsActive:is_active}
	ret_map := map[string]interface{}{}
	id, err := o.Insert(&role_data)
	if err != nil {
		fmt.Println(err)
		ret_map["code"]=10001
		ret_map["msg"]="添加失败,重新操作！"
		r.Data["json"] = ret_map
	}else {
		fmt.Println(id)
		ret_map["code"]=200
		ret_map["msg"]="添加成功！"
		r.Data["json"] = ret_map
	}
	r.ServeJSON()
}

func (r *RoleController) ToRoleUser()  {
	id, err := r.GetInt("role_id")
	if err != nil {
		logs.Info("角色添加页面获取role_id失败")
	}
	fmt.Println(id)
	o := orm.NewOrm()
	role := auth.Role{}
	_ = o.QueryTable("sys_role").Filter("id", id).One(&role)

	//已绑定用户
	o.LoadRelated(&role,"User")
	var users []auth.User
	if len(role.User)>0{
		_, _ = o.QueryTable("sys_user").Filter("is_delete", 0).Filter("is_active",1).Exclude("id__in",role.User).All(&users)
	}else { //没数据
		_, _ = o.QueryTable("sys_user").Filter("is_delete", 0).Filter("is_active",1).All(&users)
	}
	//未绑定用户


	r.Data["role"]=role
	r.Data["users"] = users
	r.TplName = "auth/role-user-add.html"
}

func (r *RoleController) DoRoleUser()  {
	role_id, err := r.GetInt("role_id")
	if err != nil {
		logs.Info("获取role_id失败")
	}
	logs.Info("要添加的角色id：",role_id)

	keys := make([]string, 0)
	r.Ctx.Input.Bind(&keys, "user_ids")
	fmt.Println(keys, len(keys))

	o := orm.NewOrm()
	role := auth.Role{Id:role_id}

	//查询删除已绑定数据
	m2m := o.QueryM2M(&role,"User")
	_, _ = m2m.Clear()

	is_ok := true
	ret_map := map[string]interface{}{}
	for _,user_id := range keys {
		fmt.Println(user_id)
		user := auth.User{Id:utils.StrToInt(user_id)}
		m2m :=o.QueryM2M(&role,"User")
		_, err := m2m.Add(user)
		if err != nil {
			is_ok = false
			ret_map["code"]=10001
			ret_map["msg"]="角色分配失败！"
			r.Data["json"] = ret_map
			break
		}
	}
	fmt.Println("+++++++++++++++",is_ok)
	if is_ok {
		ret_map["code"] = 200
		ret_map["msg"] = "角色分配成功！"
		r.Data["json"] = ret_map
	}
	fmt.Println("+++++++++++++++",ret_map)
	r.ServeJSON()
}

func (r *RoleController) ToRoleAuth()  {
	id, err := r.GetInt("role_id")
	if err != nil {
		logs.Info("权限添加页面获取role_id失败")
	}
	fmt.Println(id)
	o := orm.NewOrm()
	var role auth.Role
	_ = o.QueryTable("sys_role").Filter("id",id).One(&role)
	fmt.Println(role)
	r.Data["role"] = role
	r.TplName = "auth/role-auth-add.html"
}

func (r *RoleController) GetAuthJson()  {
	id, err := r.GetInt("role_id")
	if err != nil {
		logs.Info("权限添加页面获取role_id失败")
	}
	fmt.Println(id)
	o := orm.NewOrm()
	qs := o.QueryTable("sys_auth")
	//已绑定的权限
	role := auth.Role{Id:id}
	_, _ = o.LoadRelated(&role, "Auth")

	auth_ids_has := []int{}
	for _,auth_data := range role.Auth {
		auth_ids_has = append(auth_ids_has, auth_data.Id)
	}

	//所有权限

	var auths []auth.Auth
	_, err = qs.Filter("is_delete", 0).All(&auths)
	if err != nil {
		logs.Info("获取Auth节点失败：",err)
	}
	//var trees []auth.Tree
	//for _,auth_data := range auths {
	//	pid := auth_data.Id
	//	tree_data := auth.Tree{Id:auth_data.Id,AuthName:auth_data.AuthName,UrlFor:auth_data.UrlFor,Weight:auth_data.Weight,Children:[]*auth.Tree{}}
	//
	//	controllers.GetChildNode(pid,&tree_data)
	//	trees = append(trees, tree_data)
	//}
	//fmt.Println(trees)
	//for _,tree := range trees{
	//	for _,node := range tree.Children {
	//		fmt.Println(node)
	//	}
	//}
	auth_arr_map := []map[string]interface{}{}
	for _,auth_data := range auths {
		id := auth_data.Id
		pId := auth_data.Pid
		name := auth_data.AuthName
		if pId == 0{
			auth_map := map[string]interface{}{"id":id,"pId":pId,"name":name,"open": false}
			auth_arr_map = append(auth_arr_map, auth_map)
		}else {
			auth_map := map[string]interface{}{"id": id, "pId": pId, "name": name}
			auth_arr_map = append(auth_arr_map, auth_map)
		}
	}
	auth_maps := map[string]interface{}{}
	auth_maps["auth_arr_map"]  = auth_arr_map
	auth_maps["auth_ids_has"]  = auth_ids_has
	r.Data["json"] = auth_maps
	r.ServeJSON()
}

func (r *RoleController) DoRoleAuth()  {
	auth_ids := r.GetString("auth_ids")
	role_id,err := r.GetInt("role_id")
	if err != nil {
		logs.Info("DoRoleAuth函数获取role_id失败：",err)
	}
	//new_ids := auth_ids[1:len(auth_ids)-1]
	id_arr := strings.Split(auth_ids,",")
	o := orm.NewOrm()
	role := auth.Role{Id:role_id}
	m2m := o.QueryM2M(&role,"Auth")
	_, _ = m2m.Clear()
	is_ok := true
	ret_map := map[string]interface{}{}
	for _,id := range id_arr {
		fmt.Println(id)
		auth_id_int := utils.StrToInt(id)
		auth_data := auth.Auth{Id:auth_id_int}
		m2m :=o.QueryM2M(&role,"Auth")
		_, err := m2m.Add(auth_data)
		if err != nil {
			is_ok = false
			ret_map["code"]=10001
			ret_map["msg"]="添加失败！"
			r.Data["json"] = ret_map
			break
		}
	}
	fmt.Println("+++++++++++++++",is_ok)
	if is_ok {
		ret_map["code"] = 200
		ret_map["msg"] = "添加成功！"
		r.Data["json"] = ret_map
	}
	fmt.Println("+++++++++++++++",ret_map)
	r.ServeJSON()
}