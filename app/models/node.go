package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

const (
	NODE_IS_DELETE_FALSE = 0;
	NODE_IS_DELETE_TRUE = 1;
)

type Node struct {
	Id int64 `orm:"auto;pk;column(node_id);"`
	Ip string
	ManagerPort int
	MessagePort int
	Token string
	Comment string
	IsDelete int8
	CreateTime int64
	UpdateTime int64
}

//table name
func (node *Node) TableName() string {
	return TableName("node");
}

//插入一条节点信息
func InsertNode(node *Node) (int64, error){

	if(node.Ip == "") {
		return 0, fmt.Errorf("%s", "ip 不能为空!");
	}
	if(node.ManagerPort == 0) {
		return 0, fmt.Errorf("%s", "管理端口不能为空!");
	}
	if(node.MessagePort == 0) {
		return 0, fmt.Errorf("%s", "消息端口不能为空!");
	}
	if(node.Token == "") {
		return 0, fmt.Errorf("%s", "token 不能为空!");
	}
	if(node.Comment == "") {
		return 0, fmt.Errorf("%s", "必须填写备注!");
	}

	node.IsDelete = NODE_IS_DELETE_FALSE;
	node.CreateTime = time.Now().Unix();
	node.UpdateTime = time.Now().Unix();

	return orm.NewOrm().Insert(node);
}

//获取所有的 node
func GetNodes() ([]*Node) {
	nodes := make([]*Node, 0)
	query := orm.NewOrm().QueryTable(TableName("node"))
	query.All(&nodes);
	return nodes;
}

//根据 node_id 获取 node
func GetNodeByNodeId(nodeId int) ([]*Node)  {
	nodes := make([]*Node, 0)
	query := orm.NewOrm().QueryTable(TableName("node")).Filter("node_id", nodeId);
	query.All(&nodes);
	return nodes;
}

//修改节点信息
func UpdateNode(node *Node, fields ...string) (int64, error) {
	return orm.NewOrm().Update(node, fields...);
}

//删除节点
func DeleteNode(node *Node) (int64, error) {
	return orm.NewOrm().Delete(node);
}