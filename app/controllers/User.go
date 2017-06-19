package controllers

import (
	"wmq-admin/app/models"
	"strings"
);

type UserController struct {
	TemplateController
}

//用户列表
func (this *UserController) List() {
	users := models.GetUsers();
	this.Data["users"] = users;
	this.display("user/list");
}

//添加用户
func (this *UserController) Add()  {
	this.display("user/form");
}

//保存用户
func (this *UserController) Save() {
	user := new(models.User);

	user.Name = strings.TrimSpace(this.GetString("name"));
	user.Email = strings.TrimSpace(this.GetString("email"));
	user.Password = strings.TrimSpace(this.GetString("password"));

	userId, err := models.InsertUser(user);
	if(userId == 0 || err != nil) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("添加用户成功", "/user/list");
}