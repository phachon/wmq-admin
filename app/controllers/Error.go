package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

//错误处理
//func (this *ErrorController) redirect(code int, message string) {
//	//this.Data["message"] = message;
//	//this.TplName = code + ".html";
//}