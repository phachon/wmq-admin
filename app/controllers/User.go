package controllers

import (
	"wmq-admin/app/models"
)

type UserController struct {
	TemplateController
}

//用户列表
func (this *UserController) List() {
	users := models.GetUsers();
	this.Data["users"] = users;
	this.display("user/list.html");
}

//添加用户
func (this *UserController) Add()  {
	this.display("user/form.html");
}

//保存用户
func (this *UserController) Save() {

}