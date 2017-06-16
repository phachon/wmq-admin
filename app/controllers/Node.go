package controllers

import (
	"wmq-admin/app/models"
	"strings"
)

type NodeController struct {
	TemplateController
}

//节点列表
func (this *NodeController) List() {

	nodes := models.GetNodes();
	this.Data["nodes"] = nodes;
	this.display("node/list.html");
}

//添加节点
func (this *NodeController) Add()  {
	this.display("node/form.html");
}

//保存节点
func (this *NodeController) Save() {
	Node := new(models.Node);

	Node.Ip = strings.TrimSpace(this.GetString("ip"));
	Node.Port, _ = this.GetInt("port");
	Node.Token = strings.TrimSpace(this.GetString("token"));
	Node.Comment = strings.TrimSpace(this.GetString("comment"));

	nodeId, err := models.InsertNode(Node);
	if(nodeId == 0 || err != nil) {
		this.jsonError(err.Error(), "", make(map[string]string));
	}

	this.jsonSuccess("添加节点成功", "/node/list", make(map[string]string));
}