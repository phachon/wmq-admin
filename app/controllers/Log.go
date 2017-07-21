package controllers

import "wmq-admin/app/models"

type LogController struct {
	TemplateController
}

// 日志列表
func (this *LogController) List()  {

	nodeId, _ := this.GetInt("node_id", 1);
	keyword := this.GetString("keyword", "");
	logType := this.GetString("type", "error");

	nodes := models.GetNodes();

	models.LogSearch(nodeId, keyword, logType)

	this.Data["nodes"] = nodes;
	this.Data["node_id"] = nodeId;
	this.Data["log_type"] = logType;
	this.Data["keyword"] = keyword;
	this.display("log/list");
}