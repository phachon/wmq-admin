package controllers

type LogController struct {
	TemplateController
}

// 日志列表
func (this *LogController) List()  {
	this.display("log/list.html");
}