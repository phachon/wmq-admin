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

	messages := models.GetMessagesByNodeId(nodeId);

	nodes := models.GetNodes();
	this.Data["nodes"] = nodes;
	this.Data["node_id"] = nodeId;
	this.Data["messages"] = messages;
	this.display("message/list");
}

//添加消息
func (this *MessageController) Add()  {

	nodeId, _ := this.GetInt("node_id");

	this.layoutHtml = "layout/template";
	this.Data["node_id"] = nodeId;
	this.display("message/form");
}

//添加消息保存
func (this *MessageController) Save() {

	message := new(models.Message)

	nodeId, _ := this.GetInt("node_id");
	message.Name = this.GetString("name");
	message.Mode = this.GetString("mode");
	message.Durable, _ = this.GetBool("durable");
	message.IsNeedToken, _ = this.GetBool("is_need_token");
	message.Token = this.GetString("token");
	message.Comment = this.GetString("comment");

	res, err := models.AddMessageByNodeId(nodeId, message);
	if(!res) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("添加消息成功", "/message/list");
}

//修改消息
func (this *MessageController) Edit() {

	nodeId, _ := this.GetInt("node_id");
	messageName := this.GetString("message");

	messages := models.GetMessagesByNodeId(nodeId);

	var messageValue models.Message;
	for _, message := range messages {
		if(message.Name != messageName) {
			continue;
		}
		messageValue = message;
	}

	this.layoutHtml = "layout/template";
	this.Data["node_id"] = nodeId;
	this.Data["message"] = messageValue;
	this.display("message/edit");
}

//修改保存
func (this *MessageController) Modify()  {

	message := new(models.Message)

	nodeId, _ := this.GetInt("node_id", 0);
	message.Name = this.GetString("name", "");
	message.Mode = this.GetString("mode", "");
	message.Durable, _ = this.GetBool("durable");
	message.IsNeedToken, _ = this.GetBool("is_need_token");
	message.Token = this.GetString("token");
	message.Comment = this.GetString("comment");

	res, err := models.UpdateMessage(nodeId, message);
	if(!res) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("修改消息成功", "/message/list");
}

//删除消息
func (this *MessageController) Delete() {

	nodeId, _ := this.GetInt("node_id", 0);
	messageName := this.GetString("message", "");

	res, err := models.DeleteMessage(nodeId, messageName);
	if(!res) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("删除消息成功", "/message/list");
}

//测试发送
func (this *MessageController) Test() {

	nodeId, _ := this.GetInt("node_id", 0);

	messages := models.GetMessagesByNodeId(nodeId)

	this.layoutHtml = "layout/template";
	this.Data["node_id"] = nodeId;
	this.Data["messages"] = messages;
	this.display("message/test");
}

//发送数据
func (this *MessageController) Send() {

	nodeId, _ := this.GetInt("node_id", 0);
	message := this.GetString("message", "");
	data := this.GetString("data", "");
	routeKey := this.GetString("route_key", "");

	res, err := models.PublishMessage(nodeId, message, data, routeKey);
	if(!res) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("发送ok", "/message/list");
}