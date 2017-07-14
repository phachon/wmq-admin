package controllers

import (
	"wmq-admin/app/models"
	"strings"
	"time"
)

type NodeController struct {
	TemplateController
}

//节点列表
func (this *NodeController) List() {

	nodes := models.GetNodes();
	this.Data["nodes"] = nodes;
	this.display("node/list");
}

//添加节点
func (this *NodeController) Add()  {
	this.display("node/form");
}

//保存节点
func (this *NodeController) Save() {
	Node := new(models.Node);

	Node.Ip = strings.TrimSpace(this.GetString("ip"));
	Node.ManagerPort, _ = this.GetInt("manager_port");
	Node.MessagePort, _ = this.GetInt("message_port");
	Node.Token = strings.TrimSpace(this.GetString("token"));
	Node.Comment = strings.TrimSpace(this.GetString("comment"));

	nodeId, err := models.InsertNode(Node);
	if(nodeId == 0 || err != nil) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("添加节点成功", "/node/list");
}

//修改节点
func (this *NodeController) Edit() {
	nodeId, _ := this.GetInt("node_id");

	if(nodeId == 0) {
		this.redirect("node/list");
	}

	nodes := models.GetNodeByNodeId(nodeId);
	if(len(nodes) == 0) {
		this.redirect("node/list");
	}
	this.layoutHtml = "layout/template";

	this.Data["node"] = nodes[0];
	this.display("node/edit");
}

//修改保存
func (this *NodeController) Modify() {

	Node := new(models.Node);

	Node.Id, _ = this.GetInt64("node_id");
	Node.Ip = strings.TrimSpace(this.GetString("ip"));
	Node.ManagerPort, _ = this.GetInt("manager_port");
	Node.MessagePort, _ = this.GetInt("message_port");
	Node.Token = strings.TrimSpace(this.GetString("token"));
	Node.Comment = strings.TrimSpace(this.GetString("comment"));
	Node.UpdateTime = time.Now().Unix();

	nodeId, err := models.UpdateNode(Node);
	if(nodeId == 0 || err != nil) {
		this.jsonError(err.Error(), "");
	}

	this.jsonSuccess("修改节点成功", "/node/list");
}