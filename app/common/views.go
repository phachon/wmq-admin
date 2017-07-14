package common

import "github.com/astaxie/beego"

type Views struct {}

//加载模板函数
func (views *Views) TemplateFunc() {
	date := new(Date);
	beego.AddFuncMap("dateFormat", date.Format)
}
