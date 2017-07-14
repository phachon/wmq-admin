package models

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"strconv"
	"encoding/json"
)

type Message struct {
	Consumers   []Consumer
	Durable     bool
	IsNeedToken bool
	Mode        string
	Name        string
	Token       string
	Comment     string
}
type Consumer struct {
	ID        string
	URL       string
	RouteKey  string
	Timeout   float64
	Code      float64
	CheckCode bool
	Comment   string
}

//根据 node_id 获取 所有的消息
func GetMessagesByNodeId(nodeId int) ([]Message) {

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := selectNode.Token;
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + "/config?api-token=" + token;
	fmt.Println(nodeUrl)

	var results map[string]interface{};
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &results)
	data := results["data"].(string)

	var messages []Message
	json.Unmarshal([]byte(data), &messages)

	return messages;
}