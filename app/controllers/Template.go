package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"wmq-admin/app/common"
)

var viewPostfix = ".html";

type TemplateController struct {
	beego.Controller
	controllerName string
	methodName string
	layoutHtml string
	tplHtml string
	jsonValue
	userName string
}

type jsonValue struct {
	code int
	message string
	redirect string
	data interface{}
}

//执行前
func (this *TemplateController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.methodName = actionName

	isLogin := this.isLogin()
	if(!isLogin) {
		this.redirect("/author/index")
	}
	//判断是否登录
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

//验证登录
func (this *TemplateController) isLogin() bool {
	return true;
	//忽略 /author /error
	if(this.controllerName == "author" || this.controllerName == "error") {
		return true;
	}
	passport := beego.AppConfig.String("author.passport")
	cookie := this.Ctx.GetCookie(passport)
	//cookie 失效
	if(cookie == "") {
		return false
	}
	user := this.GetSession("author")
	//session 失效
	if(user == nil) {
		return false
	}
	encrypt := new(common.Encrypt)
	cookieValue, _ := encrypt.Base64Decode(cookie)
	if(cookieValue == "") {
		return false
	}

	identifyList := strings.Split(cookieValue, "@")
	name := identifyList[0]
	identify := identifyList[1]
	userValue := user.(map[string]interface{})

	//对比cookie 和 session name
	if(name != userValue["name"].(string)) {
		return false
	}
	//对比客户端UAG and IP
	if(identify != encrypt.Md5Encode(this.Ctx.Request.UserAgent() + this.getClientIp() + userValue["password"].(string))) {
		return false
	}

	this.userName = name;
	//success
	return true;
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
	this.Data["userName"] = this.userName;
	this.TplName = tpl + viewPostfix;
}

//成功输出json
func (this *TemplateController) jsonSuccess(message string, redirect string) {
	this.jsonValue.code = 1;
	this.jsonValue.message = message;
	this.jsonValue.redirect = redirect;
	this.jsonResult();
}

//错误输出json
func (this *TemplateController) jsonError(message string, redirect string) {
	this.jsonValue.code = 0;
	this.jsonValue.message = message;
	this.jsonValue.redirect = redirect;
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
func (this *TemplateController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

//是否是 post 请求
func (this *TemplateController) isPost() bool {
	return this.Ctx.Request.Method == "POST";
}

//获取用户IP地址
func (this *TemplateController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}