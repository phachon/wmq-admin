package main

import (
	_ "wmq-admin/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}

