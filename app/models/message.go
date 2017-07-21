package models

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"wmq-admin/app/common"
	"strconv"
)

const (
	PUBLISH_MESSAGE_PATH string = "/";
	ADD_MESSAGE_PATH string = "/message/add";
	UPDATE_MESSAGE_PATH string = "/message/update";
	DELETE_MESSAGE_PATH string = "/message/delete";
	MESSAGE_CONFIG_PATH string = "/config";

	ADD_CONSUMER_PATH string = "/consumer/add";
	UPDATE_CONSUMER_PATH string = "/consumer/update";
	DELETE_CONSUMER_PATH string = "/consumer/delete";

	RESTART_SERVICE_PATH string = "/restart";
	RELOAD_SERVICE_PATH string = "/reload";

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

type Response struct {
	Code int8
	Data []Message
}

//根据 node_id 获取 所有的消息
func GetMessagesByNodeId(nodeId int) ([]Message) {

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := selectNode.Token;
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + MESSAGE_CONFIG_PATH +"?api-token=" + token;
	fmt.Println(nodeUrl)

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);

	return res.Data;
}

// 添加一条 message 到节点
func AddMessageByNodeId(nodeId int, message *Message) (bool) {

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := selectNode.Token;
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + ADD_MESSAGE_PATH +"?api-token=" + token;
	fmt.Println(nodeUrl)

	convert := new(common.Convert)
	durable := convert.IntToTenString(convert.BoolToInt(message.Durable));
	isNeedToken := convert.IntToTenString(convert.BoolToInt(message.IsNeedToken));

	nodeUrl += "&Name=" + message.Name +
		"&Comment=" + message.Comment +
		"&Durable=" + durable +
		"&IsNeedToken=" + isNeedToken +
		"&Mode=" + message.Mode +
		"&Token=" + message.Token;

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;
	if (code == 1) {
		return true
	}else {
		return false
	}
}

//更新一条 message
func UpdateMessage(nodeId int, message *Message) (bool, error) {

	if(nodeId == 0) {
		return false, fmt.Errorf("%s", "node_id error!");
	}
	if(message.Name == "") {
		return false, fmt.Errorf("%s", "消息名称不能为空!");
	}
	if(message.Mode == "") {
		return false, fmt.Errorf("%s", "没有选择消息模式!");
	}
	if(message.IsNeedToken) {
		if(message.Token == "") {
			return false, fmt.Errorf("%s", "没有填写token!");
		}
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := selectNode.Token;
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + UPDATE_MESSAGE_PATH + "?api-token=" + token;

	convert := new(common.Convert)
	durable := convert.IntToTenString(convert.BoolToInt(message.Durable));
	isNeedToken := convert.IntToTenString(convert.BoolToInt(message.IsNeedToken));

	nodeUrl += "&Name=" + message.Name +
		"&Comment=" + message.Comment +
		"&Durable=" + durable +
		"&IsNeedToken=" + isNeedToken +
		"&Mode=" + message.Mode +
		"&Token=" + message.Token;
	fmt.Println(nodeUrl);

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;

	if (code != 1) {
		return false, fmt.Errorf("%s", "接口调用失败!");
	}

	return true, fmt.Errorf("%s", "");
}

// 删除一条 message
func DeletMessage(nodeId int, name string) bool {

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := selectNode.Token;
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + DELETE_MESSAGE_PATH + "?api-token=" + token;
	fmt.Println(nodeUrl)

	nodeUrl += "&Name=" + name;

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;
	if (code == 1) {
		return true
	}else {
		return false
	}
}

// 添加一条 consumer
func AddConsumer(nodeId int, message string, consumer *Consumer) (bool, error) {

	if(message == "") {
		return false, fmt.Errorf("%s", "没有选择消息!");
	}
	if(consumer.URL == "") {
		return false, fmt.Errorf("%s", "消费地址不能为空!");
	}
	if(consumer.Timeout == 0) {
		return false, fmt.Errorf("%s", "超时时间不能为空!");
	}
	if(consumer.RouteKey == "") {
		return false, fmt.Errorf("%s", "RouteKey 不能为空!");
	}
	consumer.Code = float64(200)
	consumer.CheckCode = false;

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := selectNode.Token;
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + ADD_CONSUMER_PATH + "?api-token=" + token;
	fmt.Println(nodeUrl)

	convert := new(common.Convert);

	nodeUrl += "&Name=" + message +
		"&URL=" + consumer.URL +
		"&Timeout=" + convert.FloatToString(consumer.Timeout, 'f', 0, 64) +
		"&Code=" + convert.FloatToString(consumer.Code, 'f', 0, 64) +
		"&CheckCode=" + convert.IntToTenString(convert.BoolToInt(consumer.CheckCode)) +
		"&Comment=" + consumer.Comment +
		"&RouteKey=" + consumer.RouteKey;

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;

	if (code != 1) {
		return false, fmt.Errorf("%s", "调用接口失败!");
	}

	return true, fmt.Errorf("%s", "");
}

// 修改一条 consumer
func UpdateConsumer(nodeId int, message string, consumer *Consumer) (bool, error) {

	if(message == "") {
		return false, fmt.Errorf("%s", "没有选择消息!");
	}
	if(consumer.URL == "") {
		return false, fmt.Errorf("%s", "消费地址不能为空!");
	}
	if(consumer.Timeout == 0) {
		return false, fmt.Errorf("%s", "超时时间不能为空!");
	}
	if(consumer.RouteKey == "") {
		return false, fmt.Errorf("%s", "RouteKey 不能为空!");
	}
	consumer.Code = float64(200)
	consumer.CheckCode = false;

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := selectNode.Token;
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + UPDATE_CONSUMER_PATH + "?api-token=" + token;
	convert := new(common.Convert);

	nodeUrl += "&Name=" + message +
		"&ID=" + consumer.ID +
		"&URL=" + consumer.URL +
		"&Timeout=" + convert.FloatToString(consumer.Timeout, 'f', 0, 64) +
		"&Code=" + convert.FloatToString(consumer.Code, 'f', 0, 64) +
		"&CheckCode=" + convert.IntToTenString(convert.BoolToInt(consumer.CheckCode)) +
		"&Comment=" + consumer.Comment +
		"&RouteKey=" + consumer.RouteKey;
	fmt.Println(nodeUrl);

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;

	if (code != 1) {
		return false, fmt.Errorf("%s", "接口调用失败!")
	}

	return true, fmt.Errorf("%s", "");
}

//删除一条 consumer
func DeleteConsumer(nodeId int, message string, consumerId string) bool {
	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := selectNode.Token;
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + DELETE_CONSUMER_PATH + "?api-token=" + token;
	fmt.Println(nodeUrl)

	nodeUrl += "&Name=" + message +
		"&ID=" + consumerId;

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;

	if (code == 1) {
		return true
	}else {
		return false
	}
}

func RestartService(nodeId int) {

}