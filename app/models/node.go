package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

const (
	STATUS_NORMAL = 0;
	STATUS_DELETE = -1;
)

type Node struct {
	Id int64 `orm:"auto;pk;column(node_id);"`
	Ip string
	ManagerPort int
	MessagePort int
	Token string
	Comment string
	Status int8
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
		return 0, fmt.Errorf("ip 不能为空!");
	}
	if(node.ManagerPort == 0) {
		return 0, fmt.Errorf("管理端口不能为空!");
	}
	if(node.MessagePort == 0) {
		return 0, fmt.Errorf("消息端口不能为空!");
	}
	if(node.Token == "") {
		return 0, fmt.Errorf("token 不能为空!");
	}
	if(node.Comment == "") {
		return 0, fmt.Errorf("必须填写备注!");
	}

	node.Status = 0
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