package controllers

import (
	"wmq-admin/app/models"
)

type MessageController struct {
	TemplateController
}

//消息列表
func (this *MessageController) List() {

	nodeId, _ := this.GetInt("node_id");
	if(nodeId == 0) {
		nodeId = 1;
	}

	messages := models.GetMessagesByNodeId(nodeId)

	nodes := models.GetNodes();
	this.Data["nodes"] = nodes;
	//this.Data["selectNode"] = selectNode;
	this.Data["messages"] = messages;
	this.display("message/list");
}

//添加消息
func (this *MessageController) Add()  {

	nodes := models.GetNodes();
	this.Data["nodes"] = nodes;
	this.display("message/form");
}

