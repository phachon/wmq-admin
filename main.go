package main

import (
	_ "wmq-admin/app/routers"
	"github.com/astaxie/beego"
	"os"
	"wmq-admin/app/models"
	"wmq-admin/app/common"
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

	if beego.AppConfig.String("runmode") == "development" {
		beego.SetLevel(beego.LevelDebug)
	} else {
		beego.SetLevel(beego.LevelInformational)
		beego.SetLogger("file", `{"filename":"`+beego.AppConfig.String("log.log_file")+`"}`)
		beego.BeeLogger.DelLogger("console")
	}

	//开启 session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//session 过期时间
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600

	//加载自定义模板函数
	new(common.Views).TemplateFunc();
	//models 初始化
	models.Init();

	//beego run
	beego.Run();
}