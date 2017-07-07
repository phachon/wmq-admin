package controllers

import (
	"wmq-admin/app/models"
	"strings"
	"time"
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

//修改用户
func (this *UserController) Edit() {

	userId, _ := this.GetInt("user_id");

	if(userId == 0) {
		this.redirect("user/list");
	}

	users := models.GetUserByUserId(userId);
	if(len(users) == 0) {
		this.redirect("user/list");
	}
	this.layoutHtml = "layout/template";

	this.Data["user"] = users[0];
	this.display("user/edit");
}

//修改保存用户
func (this *UserController) Modify() {
	user := new(models.User);

	user.Id, _ = this.GetInt64("user_id");
	user.Name = strings.TrimSpace(this.GetString("name"));
	user.Email = strings.TrimSpace(this.GetString("email"));
	user.Mobile = strings.TrimSpace(this.GetString("mobile"));
	user.UpdateTime = time.Now().Unix();

	userId, err := models.UpdateUser(user, "email", "mobile", "update_time");
	if(userId == 0 || err != nil) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("修改用户成功", "/user/list");
}

//屏蔽用户
func (this *UserController) Remove()  {
	user := new(models.User);

	user.Id, _ = this.GetInt64("user_id");
	user.IsDelete = models.USER_IS_DELETE_TRUE;

	userId, err := models.UpdateUser(user, "is_delete");
	if(userId == 0 || err != nil) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("屏蔽用户成功", "/user/list");
}

//恢复用户
func (this *UserController) Review()  {
	user := new(models.User);

	user.Id, _ = this.GetInt64("user_id");
	user.IsDelete = models.NODE_IS_DELETE_FALSE;

	userId, err := models.UpdateUser(user, "is_delete");
	if(userId == 0 || err != nil) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("恢复用户成功", "/user/list");
}