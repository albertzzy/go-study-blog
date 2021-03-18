package main

import (
	_ "myblog/routers"
	utils "myblog/utils"

	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
