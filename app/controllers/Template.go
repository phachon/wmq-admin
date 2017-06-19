package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

var viewPostfix = ".html";

type TemplateController struct {
	beego.Controller
	controllerName string
	methodName string
	layoutHtml string
	tplHtml string
	jsonValue
}

type jsonValue struct {
	code int
	message string
	redirect string
	data map[string]interface{}
}

//执行前
func (this *TemplateController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.methodName = actionName

	//默认layout
	this.layoutHtml = "layout/default"
	//默认tpl
	this.tplHtml = "author/login"
	//默认json返回
	this.jsonValue.code = 1;
	this.jsonValue.message = "";
	this.jsonValue.redirect = "";
	this.jsonValue.data = make(map[string]interface{});
}

//执行后
func (this *TemplateController) finish() {
	
}

//渲染模板
func (this *TemplateController) display(tpl string) {
	if(tpl != "") {
		this.tplHtml = tpl
	}
	this.Layout = this.layoutHtml + viewPostfix;
	this.Data["navName"] = this.controllerName;
	this.TplName = tpl + viewPostfix;
}

//成功输出json
func (this *TemplateController) jsonSuccess(message string, redirect string) {
	this.jsonValue.code = 1;
	this.jsonValue.message = message;
	this.jsonResult();
}

//错误输出json
func (this *TemplateController) jsonError(message string, redirect string) {
	this.jsonValue.code = 0;
	this.jsonValue.message = message;
	this.jsonResult();
}

//输出json
func (this *TemplateController) jsonResult() {
	body := make(map[string]interface{});
	body["code"] = this.jsonValue.code;
	body["message"] = this.jsonValue.message;
	body["redirect"] = this.jsonValue.redirect;
	body["data"] = this.jsonValue.data;

	this.Data["json"] = body;
	this.ServeJSON();
	this.StopRun();
}

//302跳转
func (this *TemplateController) redirect(tpl ...string) {

}

//是否是 post 请求
func (this *TemplateController) isPost() bool {
	return this.Ctx.Request.Method == "POST";
}