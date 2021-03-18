package controllers

import (
	"fmt"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	fmt.Println("=====islogin===", this.IsLogin, this.Loginuser)

	this.Data["IsLogin"] = this.IsLogin
	this.Data["Loginuser"] = this.Loginuser
	this.TplName = "home.html"
}
