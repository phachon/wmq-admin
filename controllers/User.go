package controllers

type UserController struct {
	TemplateController
}

//用户列表
func (this *UserController) List() {
	this.display("user/list.html");
}

//添加用户
func (this *UserController) Add()  {
	this.display("user/form.html");
}