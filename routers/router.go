package routers

import (
	"ions_zhiliao/controllers"
	"github.com/astaxie/beego"
	"ions_zhiliao/controllers/auth"
	"ions_zhiliao/controllers/caiwu"
	"ions_zhiliao/controllers/cars"
	"ions_zhiliao/controllers/echarts"
	"ions_zhiliao/controllers/login"
	"ions_zhiliao/controllers/news"
	"ions_zhiliao/controllers/user"
)

func init() {

	//未登录可请求
    beego.Router("/login", &login.LoginController{})
    beego.Router("/admin/log_out", &login.LoginController{},"get:LogOut")
    //修改验证码
    beego.Router("/change_captcha", &login.LoginController{},"get:ChangeCaptcha")
    //登录才能访问
    //后端首页
	beego.Router("/admin/index", &controllers.HomeController{})
	beego.Router("/admin/index/notify", &controllers.HomeController{},"get:NotifyList")
	beego.Router("/admin/index/read_notify", &controllers.HomeController{},"get:ReadNotify")
	beego.Router("/admin/welcome", &controllers.HomeController{},"get:Welcome")
    //用户模块
	beego.Router("/admin/user/list", &user.UserController{},"get:List")
    //用户添加页
	beego.Router("/admin/user/to_add", &user.UserController{},"get:ToAdd")
	beego.Router("/admin/user/do_add", &user.UserController{},"post:DoAdd")
    //修改启用状态
	beego.Router("/admin/user/is_active", &user.UserController{},"post:IsActive")
    //删除用户
	beego.Router("/admin/user/delete", &user.UserController{},"get:Delete")
    //重置密码
	beego.Router("/admin/user/reset_pwd", &user.UserController{},"get:ResetPassword")
	beego.Router("/admin/user/to_edit", &user.UserController{},"get:ToUpdate")
	//用户修改
	beego.Router("/admin/user/do_update", &user.UserController{},"post:DoUpdate")
    //批量删除
	beego.Router("/admin/user/muli_delete", &user.UserController{},"post:MuliDelete")
    //权限
	beego.Router("/admin/auth/list", &auth.AuthController{},"get:List")
	beego.Router("/admin/auth/to_add", &auth.AuthController{},"get:ToAdd")
	beego.Router("/admin/auth/do_add", &auth.AuthController{},"post:DoAdd")
    //角色
	beego.Router("/admin/role/list", &auth.RoleController{},"get:List")
	beego.Router("/admin/role/to_add", &auth.RoleController{},"get:ToAdd")
	beego.Router("/admin/role/do_add", &auth.RoleController{},"post:DoAdd")
    //用户角色
	beego.Router("/admin/role/to_role_user", &auth.RoleController{},"get:ToRoleUser")
	beego.Router("/admin/role/do_role_user", &auth.RoleController{},"post:DoRoleUser")

    //角色权限
	beego.Router("/admin/role/to_role_auth", &auth.RoleController{},"get:ToRoleAuth")
	beego.Router("/admin/role/do_role_auth", &auth.RoleController{},"post:DoRoleAuth")
	beego.Router("/admin/role/get_auth_json", &auth.RoleController{},"get:GetAuthJson")
	//个人信息
	beego.Router("/admin/user/my_center", &user.MyCenterController{})
	beego.Router("/admin/user/salary_slip_list", &user.SalarySlipController{})
	beego.Router("/admin/user/salary_slip_detail", &user.SalarySlipController{},"get:Detail")
    //财务
	beego.Router("/admin/caiwu/salary_slip_list", &caiwu.CaiWuSalarySlipController{})
	beego.Router("/admin/caiwu/to_salary_slip_import", &caiwu.CaiWuSalarySlipController{},"get:ToImportExcel")
	beego.Router("/admin/caiwu/do_salary_slip_import", &caiwu.CaiWuSalarySlipController{},"post:DoImportExcel")

	beego.Router("/admin/caiwu/echart_data_list", &caiwu.CaiWuEchartDataController{})
	beego.Router("/admin/caiwu/to_echart_data_import", &caiwu.CaiWuEchartDataController{},"get:ToImportExcel")
	beego.Router("/admin/caiwu/do_echart_data_import", &caiwu.CaiWuEchartDataController{},"post:DoImportExcel")
	//内容管理CategoryController
	beego.Router("/admin/news/category_list", &news.CategoryController{})
	beego.Router("/admin/news/to_category", &news.CategoryController{},"get:ToAdd")
	beego.Router("/admin/news/do_category", &news.CategoryController{},"post:DoAdd")
	beego.Router("/admin/news/news_list", &news.NewsController{})
	beego.Router("/admin/news/to_news", &news.NewsController{},"get:ToAdd")
	beego.Router("/admin/news/upload_img", &news.NewsController{},"post:UploadImg")
	beego.Router("/admin/news/do_news", &news.NewsController{},"post:DoAdd")
	beego.Router("/admin/news/to_edit", &news.NewsController{},"get:ToEdit")
	beego.Router("/admin/news/do_edit", &news.NewsController{},"post:DoEdit")
	//车辆管理
	beego.Router("/admin/cars/cars_brand_list", &cars.CarBrandController{})
	beego.Router("/admin/cars/to_cars_brand", &cars.CarBrandController{},"get:ToAdd")
	beego.Router("/admin/cars/do_cars_brand", &cars.CarBrandController{},"post:DoAdd")

	beego.Router("/admin/cars/cars_list", &cars.CarsController{})
	beego.Router("/admin/cars/to_cars", &cars.CarsController{},"get:ToAdd")
	beego.Router("/admin/cars/do_cars", &cars.CarsController{},"post:DoAdd")

	beego.Router("/admin/cars/cars_apply_list", &cars.CarsApplyController{})
	beego.Router("/admin/cars/to_apply", &cars.CarsApplyController{},"get:ToApply")
	beego.Router("/admin/cars/do_apply", &cars.CarsApplyController{},"post:DoApply")

	beego.Router("/admin/cars/my_apply_list", &cars.CarsApplyController{},"get:MyApply")
	beego.Router("/admin/cars/audit_apply_list", &cars.CarsApplyController{},"get:AuditApply")
	beego.Router("/admin/cars/to_audit_apply", &cars.CarsApplyController{},"get:ToAuditApply")
	beego.Router("/admin/cars/Do_audit_apply", &cars.CarsApplyController{},"post:DoAuditApply")
	beego.Router("/admin/cars/Do_return", &cars.CarsApplyController{},"get:DoReturn")

	//报表管理
	beego.Router("/admin/echarts/caiwu", &echarts.EchartsCaiwuController{})
	beego.Router("/admin/echarts/get_caiwu_data", &echarts.EchartsCaiwuController{},"get:GetCaiWuChart")

	beego.Router("/admin/echarts/course", &echarts.EchartsCourseController{})
	beego.Router("/admin/echarts/get_course_data", &echarts.EchartsCourseController{},"get:GetCourseChart")

	beego.Router("/admin/echarts/business", &echarts.EchartsBusinessController{})
	beego.Router("/admin/echarts/get_business_data", &echarts.EchartsBusinessController{},"get:GetBusinessChart")

}
