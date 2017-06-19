package routers

import (
	"wmq-admin/app/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.AuthorController{}, "*:Index")
    beego.AutoRouter(&controllers.AuthorController{})
    beego.AutoRouter(&controllers.IndexController{})
    beego.AutoRouter(&controllers.UserController{})
    beego.AutoRouter(&controllers.NodeController{})
    beego.AutoRouter(&controllers.MessageController{})
    beego.AutoRouter(&controllers.ConsumerController{})
    beego.AutoRouter(&controllers.LogController{})
    beego.AutoRouter(&controllers.ProfileController{})
}
