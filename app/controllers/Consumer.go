package controllers

type ConsumerController struct {
	TemplateController
}

//消费列表
func (this *ConsumerController) List()  {
	this.display("consumer/list.html");
}

//添加消费者
func (this *ConsumerController) Add()  {
	this.display("consumer/form.html");
}