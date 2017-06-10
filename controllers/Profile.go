package controllers

type ProfileController struct {
	TemplateController
}

//个人资料
func (this *ProfileController) Index() {
	this.display("profile/index.html");
}

//修改密码
func (this *ProfileController) Repass() {
	this.display("profile/repass.html");
}