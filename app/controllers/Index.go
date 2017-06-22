package controllers

type IndexController struct {
	TemplateController
}

//首页
func (this *IndexController) Main()  {
	this.display("index/main");
}

