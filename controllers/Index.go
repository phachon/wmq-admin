package controllers

type IndexController struct {
	TemplateController
}

//首页
func (this *IndexController) Index()  {
	this.display("index/main.html");
}

