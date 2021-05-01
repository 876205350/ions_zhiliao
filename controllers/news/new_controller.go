package news

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"ions_zhiliao/models/news"
	"ions_zhiliao/utils"
	"math"
	"strconv"
	"time"
)

type NewsController struct {
	beego.Controller
}

func (n *NewsController) Get()  {
	o := orm.NewOrm()
	var news_data []news.News
	qs := o.QueryTable("sys_news")
	//_, _ = qs.Filter("is_delete", 0).All(&news_data)
	pageSize := 10
	pageNo, err := n.GetInt("page")
	if err != nil {
		pageNo = 1
	}
	var count int64 = 0
	offsetNum := pageSize * (pageNo-1)
	kw := n.GetString("kw")
	if kw != "" {
		count, _ = qs.Filter("title__contains", kw).Filter("is_delete", 0).Count()
		_, _ = qs.Filter("title__contains",kw).Filter("is_delete", 0).Limit(pageSize).Offset(offsetNum).RelatedSel().All(&news_data)
	}else {
		count, _ = qs.Filter("is_delete", 0).Count()
		_, _ = qs.Filter("is_delete", 0).Limit(pageSize).Offset(offsetNum).RelatedSel().All(&news_data)
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

	n.Data["count"] = count
	n.Data["pageNo"] = pageNo
	n.Data["prePage"] = prePage
	n.Data["nextPage"] = nextPage
	n.Data["countPage"] = countPage
	n.Data["page_map"] = page_map
	n.Data["kw"] = kw
	n.Data["news_data"] = news_data
	n.TplName = "news/news_list.html"
}

func (n *NewsController) ToAdd()  {
	o := orm.NewOrm()
	var categories []news.Category
	qs := o.QueryTable("sys_category")
	_, _ = qs.Filter("is_delete", 0).All(&categories)
	n.Data["categories"] = categories
	n.TplName = "news/news_add.html"
}

func (n *NewsController) UploadImg()  {
	f, h, err := n.GetFile("file")
	ret_map := map[string]interface{}{}
	defer func() {
		f.Close()
	}()
	if err != nil {
		ret_map["code"] = 10001
		ret_map["msg"] = "文件上传失败"
		logs.Info("excel获取失败：",err)
		n.Data["json"] = ret_map
		n.ServeJSON()
		return
	}
	file_name :=  h.Filename
	fmt.Println(file_name)
	time_unix_int := time.Now().Unix()
	time_unix_str := strconv.FormatInt(time_unix_int,10)
	fmt.Println(time_unix_str)
	file_path := "upload/news_img/"+time_unix_str+"-"+file_name
	err = n.SaveToFile("file", file_path)
	if err != nil {
		ret_map["code"] = 10001
		ret_map["msg"] = "文件保存失败"
	}else {
		img_link := "http://127.0.0.1:8080/"+file_path
		ret_map["code"] = 200
		ret_map["msg"] = "文件保存成功"
		ret_map["link"] = img_link
	}
	n.Data["json"] = ret_map
	n.ServeJSON()
}

func (n *NewsController) DoAdd()  {
	content := n.GetString("content")
	title := n.GetString("title")
	category_id,_ := n.GetInt("category_id")
	is_active,_ := n.GetInt("is_active")

	category := news.Category{Id:category_id}
	o := orm.NewOrm()
	news_data := news.News{
		Content:content,
		Title:title,
		Category:&category,
		IsActive:is_active,
	}
	_,err := o.Insert(&news_data)

	message_map := map[string]interface{}{}
	if err != nil {
		message_map["code"] = 10001
		message_map["msg"] = "添加失败"
	}else {
	message_map["code"] = 200
	message_map["msg"] = "添加成功"
	}
	n.Data["json"] = message_map
	n.ServeJSON()
}

func (n *NewsController) ToEdit()  {
	id, _ := n.GetInt("id")
	o := orm.NewOrm()
	news_data := news.News{}
	qs := o.QueryTable("sys_news")
	_ = qs.Filter("id",id).Filter("is_delete", 0).RelatedSel().One(&news_data)

	var categories []news.Category
	 _, _ = o.QueryTable("sys_category").Filter("is_delete", 0).Exclude("id", news_data.Category.Id).All(&categories)
	n.Data["categories"] = categories
	n.Data["news_data"] = news_data
	n.TplName = "news/news_edit.html"
}

func (n *NewsController) DoEdit()  {
	id, _ := n.GetInt("news_id")
	content := n.GetString("content")
	title := n.GetString("title")
	category_id,_ := n.GetInt("category_id")
	is_active,_ := n.GetInt("is_active")
	o := orm.NewOrm()
	qs := o.QueryTable("sys_news")
	_, err := qs.Filter("id", id).Update(orm.Params{
		"title":       title,
		"content":     content,
		"category_id": category_id,
		"is_active":   is_active,
	})
	ret_map := make(map[string]interface{},5)
	if err != nil {
		ret_map["code"] = 10001
		ret_map["msg"] = "新闻编辑失败！"
	}else {
		ret_map["code"] = 200
		ret_map["msg"] = "新闻编辑成功！"
	}
	n.Data["json"] = ret_map
	n.ServeJSON()
}