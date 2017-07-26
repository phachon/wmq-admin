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

	_, logs := models.LogSearch(nodeId, keyword, logType)
	_, logDownloads := models.LogDownload(nodeId)

	this.Data["nodes"] = nodes;
	this.Data["node_id"] = nodeId;
	this.Data["log_type"] = logType;
	this.Data["keyword"] = keyword;
	this.Data["logs"] = logs;
	this.Data["logDownloads"] = logDownloads;

	this.display("log/list");
}

// 日志列表
func (this *LogController) Download()  {

	nodeId, _ := this.GetInt("node_id", 1);

	_, logDownloads := models.LogDownload(nodeId)

	this.Data["logDownloads"] = logDownloads;

	this.layoutHtml = "layout/template";

	this.display("log/download");
}