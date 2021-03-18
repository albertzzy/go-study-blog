package routers

import (
	"myblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/article/add", &controllers.WriteArticleController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.ExitController{})
	beego.Router("/register", &controllers.RegisterController{})

	//显示文章详情
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	//文件上传
	beego.Router("/upload", &controllers.UploadController{})
}
