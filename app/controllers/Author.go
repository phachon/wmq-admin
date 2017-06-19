package controllers

import (
	"strings"
	"wmq-admin/app/models"
	"wmq-admin/app/common"
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

	this.jsonSuccess("登录成功", "/index/main");
}