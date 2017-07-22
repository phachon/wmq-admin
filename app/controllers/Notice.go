package controllers

import (
	"wmq-admin/app/models"
	"strings"
)

type NoticeController struct {
	TemplateController
}

//保存公告
func (this *NoticeController) Save() {
	Notice := new(models.Notice);

	Notice.UserName = strings.TrimSpace(this.GetString("user_name"));
	Notice.Message = strings.TrimSpace(this.GetString("message"));

	nodeId, err := models.InsertNotice(Notice);
	if(nodeId == 0 || err != nil) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("添加公告成功", "/");
}