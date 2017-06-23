package controllers

import (
	"strings"
	"wmq-admin/app/models"
	"wmq-admin/app/common"
	"github.com/astaxie/beego"
)

type AuthorController struct {
	TemplateController
}

//login index
func (this *AuthorController) Index() {
	this.layoutHtml = "layout/author";
	this.display("author/login");
}

//login
func (this *TemplateController) Login()  {

	name := strings.TrimSpace(this.GetString("name"));
	password := strings.TrimSpace(this.GetString("password"));

	users := models.GetUserByName(name);

	if(len(users) == 0) {
		this.jsonError("账号错误!", "");
	}
	password = common.Md5Encode(password);

	if(users[0].Password != password) {
		this.jsonError("账号或密码错误!", "");
	}

	user := make(map[string]interface{})
	user["user_id"] = users[0].Id
	user["name"] = users[0].Name
	user["email"] = users[0].Email
	user["password"] = users[0].Password
	user["mobile"] = users[0].Mobile
	user["is_delete"] = users[0].IsDelete
	user["create_time"] = users[0].CreateTime
	user["update_time"] = users[0].UpdateTime

	//保存 session
	this.SetSession("author", user)
	//保存 cookie
	encrypt := new(common.Encrypt)
	identify := encrypt.Md5Encode(this.Ctx.Request.UserAgent() + this.getClientIp() + password)
	passportValue := encrypt.Base64Encode(name + "@" + identify)
	passport := beego.AppConfig.String("author.passport")
	this.Ctx.SetCookie(passport, passportValue, 3600)

	this.jsonSuccess("登录成功", "/index/main");
}