package main

import (
	_ "wmq-admin/app/routers"
	"github.com/astaxie/beego"
	"os"
	"wmq-admin/app/models"
)

func main() {
	//根据环境变量动态加载
	env := os.Getenv("GOENV");
	if(env == "") {
		env = "development";
	}
	if(env == "development") {
		beego.LoadAppConfig("ini", "conf/development.conf");
	}
	if(env == "testing") {
		beego.LoadAppConfig("ini", "conf/testing.conf");
	}
	if(env == "production") {
		beego.LoadAppConfig("ini", "conf/production.conf")
	}

	models.Init();
	beego.Run();
}

