package controllers

import (
	"fmt"
	"myblog/models"
	utils "myblog/utils"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:", username, ",password:", password)

	/*
		这里只是用session在不在来判断是否登录，没有用到cookie
	*/

	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)
	if id > 0 {
		this.SetSession("loginuser", username)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	this.ServeJSON()
}
