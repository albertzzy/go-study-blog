package controllers

import "fmt"

type ExitController struct {
	BaseController
}

func (this *ExitController) Get() {

	fmt.Println("IsLogin:======", this.DelSession)

	this.DelSession("loginuser")
	this.Redirect("/", 302)
}
