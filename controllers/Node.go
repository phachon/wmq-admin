package controllers

type NodeController struct {
	TemplateController
}

//用户列表
func (this *NodeController) List() {
	this.display("node/list.html");
}

//添加用户
func (this *NodeController) Add()  {
	this.display("node/form.html");
}