package controllers

import "wmq-admin/app/models"

type IndexController struct {
	TemplateController
}

//首页
func (this *IndexController) Main()  {

	users := models.GetUsers();
	this.Data["users"] = users;

	notices := models.GetNotices();
	this.Data["notices"] = notices;

	this.display("index/main");
}

