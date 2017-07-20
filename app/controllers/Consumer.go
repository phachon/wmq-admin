package controllers

import (
	"wmq-admin/app/models"
	"fmt"
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
	this.Data["messages"] = messages;

	this.display("consumer/list");
}

//添加消费者
func (this *ConsumerController) Add() {
	this.display("consumer/form");
}

//修改消费者
func (this *ConsumerController) Edit() {

	nodeId, _ := this.GetInt("node_id");
	messageName := this.GetString("message");
	consumerId := this.GetString("consumer_id");

	messages := models.GetMessagesByNodeId(nodeId);
	for _, message := range messages {
		if(message.Name != messageName) {
			continue;
		}
		for _, consumer := range message.Consumers {
			if(consumer.ID != consumerId) {
				continue;
			}
			fmt.Println(consumer)
		}
	}


	this.layoutHtml = "layout/template";
	this.display("consumer/edit");
}