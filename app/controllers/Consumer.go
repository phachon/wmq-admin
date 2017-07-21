package controllers

import (
	"wmq-admin/app/models"
)

type ConsumerController struct {
	TemplateController
}

//消费列表
func (this *ConsumerController) List()  {

	nodeId, _ := this.GetInt("node_id");
	if(nodeId == 0) {
		nodeId = 1;
	}

	messages := models.GetMessagesByNodeId(nodeId)

	nodes := models.GetNodes();
	this.Data["nodes"] = nodes;
	this.Data["node_id"] = nodeId;
	this.Data["messages"] = messages;

	this.display("consumer/list");
}

//添加消费者
func (this *ConsumerController) Add() {

	nodeId, _ := this.GetInt("node_id");

	messages := models.GetMessagesByNodeId(nodeId);

	this.layoutHtml = "layout/template";
	this.Data["node_id"] = nodeId;
	this.Data["messages"] = messages;
	this.display("consumer/form");
}

//添加保存
func (this *ConsumerController) Save() {

	var consumer = new(models.Consumer);

	nodeId, _ := this.GetInt("node_id");
	messageName := this.GetString("message");
	consumer.URL = this.GetString("url");
	consumer.RouteKey = this.GetString("route_key");
	consumer.Timeout, _ = this.GetFloat("timeout");
	consumer.Comment = this.GetString("comment");

	res, err := models.AddConsumer(nodeId, messageName, consumer);
	if(!res) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("添加消费者成功!", "/consumer/list");
}

//修改消费者
func (this *ConsumerController) Edit() {

	nodeId, _ := this.GetInt("node_id");
	messageName := this.GetString("message");
	consumerId := this.GetString("consumer_id");

	messages := models.GetMessagesByNodeId(nodeId);

	var consumerValue models.Consumer;
	for _, message := range messages {
		if(message.Name != messageName) {
			continue;
		}
		for _, consumer := range message.Consumers {
			if(consumer.ID != consumerId) {
				continue;
			}
			consumerValue = consumer;
		}
	}

	this.layoutHtml = "layout/template";
	this.Data["consumer"] = consumerValue;
	this.Data["node_id"] = nodeId;
	this.Data["selectMessage"] = messageName;
	this.display("consumer/edit");
}

//修改保存
func (this *ConsumerController) Modify() {

	var consumer = new(models.Consumer);

	nodeId, _ := this.GetInt("node_id");
	messageName := this.GetString("message");
	consumer.URL = this.GetString("url");
	consumer.RouteKey = this.GetString("route_key");
	consumer.Timeout, _ = this.GetFloat("timeout");
	consumer.Comment = this.GetString("comment");
	consumer.ID = this.GetString("consumer_id");

	res, err := models.UpdateConsumer(nodeId, messageName, consumer);
	if(!res) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("修改消费者成功!", "/consumer/list");
}

//删除消费者
func (this *ConsumerController) Delete() {

	nodeId, _ := this.GetInt("node_id");
	messageName := this.GetString("message");
	consumerId := this.GetString("consumer_id");

	res, err := models.DeleteConsumer(nodeId, messageName, consumerId);
	if(!res) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("删除消费者成功!", "/consumer/list");
}