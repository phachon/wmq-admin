package controllers

type MessageController struct {
	TemplateController
}

//消息列表
func (this *MessageController) List()  {
	this.display("message/list.html");
}

//添加消息
func (this *MessageController) Add()  {
	this.display("message/form.html");
}

