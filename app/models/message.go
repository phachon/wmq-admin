package models

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"wmq-admin/app/common"
	"strconv"
	"net/url"
	"strings"
)

const (
	PUBLISH_MESSAGE_PATH string = "/";

	ADD_MESSAGE_PATH string = "/message/add";
	UPDATE_MESSAGE_PATH string = "/message/update";
	DELETE_MESSAGE_PATH string = "/message/delete";
	MESSAGE_CONFIG_PATH string = "/config";
	MESSAGE_STATUS_PATH string = "/message/status";

	ADD_CONSUMER_PATH string = "/consumer/add";
	UPDATE_CONSUMER_PATH string = "/consumer/update";
	DELETE_CONSUMER_PATH string = "/consumer/delete";

	RESTART_SERVICE_PATH string = "/restart";
	RELOAD_SERVICE_PATH string = "/reload";

	LOG_SEARCH_PATH string = "/log";
	LOG_FILE_LIST_PATH string = "/log/list";
	LOG_FILE_DOWNLOAD string = "/log/file";
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

type WmqLog struct {
	Content               string
	Timestamp             int64
	Milliseconds          int64
	TimestampMilliseconds int64
	Level                 uint8
	LevelString           string
	Fields                map[string]string
	FieldsString          string
}

type Status struct {
	Count int
	ID string
	LastTime string
	MsgName string
}

type Response struct {
	Code int8
	Data string
}

//根据 node_id 获取 所有的消息
func GetMessagesByNodeId(nodeId int) ([]Message) {

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + MESSAGE_CONFIG_PATH +"?api-token=" + token;
	fmt.Println(nodeUrl)

	type ConfigResponse struct {
		Code int8
		Data []Message
	}

	var res ConfigResponse;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);

	return res.Data;
}

// 添加一条 message 到节点
func AddMessageByNodeId(nodeId int, message *Message) (bool, error) {

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
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + ADD_MESSAGE_PATH +"?api-token=" + token;
	fmt.Println(nodeUrl)

	convert := new(common.Convert)
	durable := convert.IntToTenString(convert.BoolToInt(message.Durable));
	isNeedToken := convert.IntToTenString(convert.BoolToInt(message.IsNeedToken));

	nodeUrl += "&Name=" + message.Name +
		"&Comment=" + url.QueryEscape(message.Comment) +
		"&Durable=" + durable +
		"&IsNeedToken=" + isNeedToken +
		"&Mode=" + message.Mode +
		"&Token=" + url.QueryEscape(message.Token);

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;
	if (code != 1) {
		return false, fmt.Errorf("%s", "调用接口失败!");
	}

	return true, fmt.Errorf("%s", "");
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
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + UPDATE_MESSAGE_PATH + "?api-token=" + token;

	convert := new(common.Convert)
	durable := convert.IntToTenString(convert.BoolToInt(message.Durable));
	isNeedToken := convert.IntToTenString(convert.BoolToInt(message.IsNeedToken));

	nodeUrl += "&Name=" + message.Name +
		"&Comment=" + url.QueryEscape(message.Comment) +
		"&Durable=" + durable +
		"&IsNeedToken=" + isNeedToken +
		"&Mode=" + message.Mode +
		"&Token=" + url.QueryEscape(message.Token);
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
func DeleteMessage(nodeId int, name string) (bool, error) {

	if(nodeId == 0) {
		return false, fmt.Errorf("%s", "node_id error!");
	}
	if(name == "") {
		return false, fmt.Errorf("%s", "message name error!");
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + DELETE_MESSAGE_PATH + "?api-token=" + token;
	fmt.Println(nodeUrl)

	nodeUrl += "&Name=" + name;

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;
	if (code != 1) {
		return false, fmt.Errorf("%s", "调用接口失败!");
	}

	return true, fmt.Errorf("%s", "");
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
	if(consumer.CheckCode) {
		if(consumer.Code == 0) {
			return false, fmt.Errorf("%s", "检查code码不能为空!");
		}
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + ADD_CONSUMER_PATH + "?api-token=" + token;

	convert := new(common.Convert);

	nodeUrl += "&Name=" + message +
		"&URL=" + url.QueryEscape(consumer.URL) +
		"&Timeout=" + convert.FloatToString(consumer.Timeout, 'f', 0, 64) +
		"&Code=" + convert.FloatToString(consumer.Code, 'f', 0, 64) +
		"&CheckCode=" + convert.IntToTenString(convert.BoolToInt(consumer.CheckCode)) +
		"&Comment=" + url.QueryEscape(consumer.Comment) +
		"&RouteKey=" + url.QueryEscape(consumer.RouteKey);
	fmt.Println(nodeUrl);

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
	if(consumer.CheckCode) {
		if(consumer.Code == 0) {
			return false, fmt.Errorf("%s", "检查code码不能为空!");
		}
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + UPDATE_CONSUMER_PATH + "?api-token=" + token;
	convert := new(common.Convert);

	nodeUrl += "&Name=" + message +
		"&ID=" + consumer.ID +
		"&URL=" + url.QueryEscape(consumer.URL) +
		"&Timeout=" + convert.FloatToString(consumer.Timeout, 'f', 0, 64) +
		"&Code=" + convert.FloatToString(consumer.Code, 'f', 0, 64) +
		"&CheckCode=" + convert.IntToTenString(convert.BoolToInt(consumer.CheckCode)) +
		"&Comment=" + url.QueryEscape(consumer.Comment) +
		"&RouteKey=" + url.QueryEscape(consumer.RouteKey);
	fmt.Println(nodeUrl);

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;

	if (code != 1) {
		return false, fmt.Errorf("%s", "接口调用失败:" + res.Data)
	}

	return true, fmt.Errorf("%s", "");
}

//删除一条 consumer
func DeleteConsumer(nodeId int, message string, consumerId string) (bool, error) {

	if(nodeId == 0) {
		return false, fmt.Errorf("%s", "node_id error!");
	}
	if(message == "") {
		return false, fmt.Errorf("%s", "message error!");
	}
	if(consumerId == "") {
		return false, fmt.Errorf("%s", "consumer_id error!");
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + DELETE_CONSUMER_PATH + "?api-token=" + token;

	nodeUrl += "&Name=" + message + "&ID=" + consumerId;
	fmt.Println(nodeUrl);

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;

	if (code != 1) {
		return false, fmt.Errorf("%s", "调用接口失败!");
	}

	return true, fmt.Errorf("%s", "");
}

//重启服务
func RestartService(nodeId int) (bool, error) {

	if(nodeId == 0) {
		return false, fmt.Errorf("%s", "node_id error!");
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + RESTART_SERVICE_PATH + "?api-token=" + token;

	fmt.Println(nodeUrl);

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;

	if (code != 1) {
		return false, fmt.Errorf("%s", "调用接口失败:" + res.Data);
	}

	return true, fmt.Errorf("%s", "");
}

//重载服务
func ReloadService(nodeId int) (bool, error) {

	if(nodeId == 0) {
		return false, fmt.Errorf("%s", "node_id error!");
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + RELOAD_SERVICE_PATH + "?api-token=" + token;

	fmt.Println(nodeUrl);

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;

	if (code != 1) {
		return false, fmt.Errorf("%s", "调用接口失败:" + res.Data);
	}

	return true, fmt.Errorf("%s", "");
}

//获取 node_id 所有消费者状态
func ConsumerStatus(nodeId int) ([]map[string]interface{}, error){

	if(nodeId == 0) {
		return nil, fmt.Errorf("%s", "node_id error!");
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + MESSAGE_STATUS_PATH + "?api-token=" + token;

	type StatusResponse struct {
		Code int
		Data []map[string]interface{};
	}

	var consumerStatus []map[string]interface{};

	messages := GetMessagesByNodeId(nodeId);
	for _, message := range messages {
		statusUrl := nodeUrl + "&Name=" + message.Name;
		fmt.Println(statusUrl);
		var res StatusResponse;
		response, _ := httplib.Get(statusUrl).String();
		json.Unmarshal([]byte(response), &res);
		if(res.Code == 1) {
			consumerStatus = append(consumerStatus, res.Data...);
		}
	}

	return consumerStatus, nil;
}

//publish a message
func PublishMessage(nodeId int, messageName string, data string, routeKey string) (bool, error) {

	if(nodeId == 0) {
		return false, fmt.Errorf("%s", "node_id error!");
	}
	if(messageName == "") {
		return false, fmt.Errorf("%s", "没有选择 message!");
	}
	if(data == "") {
		return false, fmt.Errorf("%s", "data is empty!");
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	messagePort := selectNode.MessagePort;
	publishUrl := "http://" + ip + ":" + strconv.Itoa(messagePort) + PUBLISH_MESSAGE_PATH + messageName + "?" + data;

	fmt.Println("Test Publish Message: " + publishUrl);

	messages := GetMessagesByNodeId(nodeId)
	var messageValue Message;
	for _, message := range messages {
		if(message.Name != messageName) {
			continue;
		}
		messageValue = message;
	}

	request := httplib.Get(publishUrl);
	if(messageValue.IsNeedToken) {
		request.Header("Token", messageValue.Token);
	}
	if(routeKey != "") {
		request.Header("RouteKey", routeKey);
	}

	request.Response();

	return true, nil;
}

//搜索日志
func LogSearch(nodeId int, keyword string, logType string) (error, []WmqLog) {

	if(nodeId == 0) {
		return fmt.Errorf("%s", "node_id error!"), nil;
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + LOG_SEARCH_PATH + "?api-token=" + token;
	nodeUrl += "&keyword=" + keyword + "&type=" + logType

	fmt.Println(nodeUrl);

	var res Response;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;

	if (code != 1) {
		return fmt.Errorf("%s", "调用接口失败:" + res.Data), nil;
	}

	var wmqLog WmqLog
	var logResults []WmqLog

	data := strings.Split(res.Data, "\n")
	for _,log := range data {
		json.Unmarshal([]byte(log), &wmqLog)
		logResults = append(logResults, wmqLog)
	}

	return nil, logResults;
}

//搜索日志
func LogDownload(nodeId int) (error, map[string]string) {

	if(nodeId == 0) {
		return fmt.Errorf("%s", "node_id error!"), nil;
	}

	selectNode := GetNodeByNodeId(nodeId)[0];
	ip := selectNode.Ip;
	managerPort := selectNode.ManagerPort;
	token := url.QueryEscape(selectNode.Token);
	nodeUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + LOG_FILE_LIST_PATH + "?api-token=" + token;

	fmt.Println(nodeUrl);

	type LogFIle struct {
		Code int
		Data []string
	}

	var res LogFIle;
	response, _ := httplib.Get(nodeUrl).String();
	json.Unmarshal([]byte(response), &res);
	code := res.Code;

	if(code != 1) {
		return fmt.Errorf("%s", "调用接口失败:"), nil;
	}

	downloadInfo := make(map[string]string)

	downloadUrl := "http://" + ip + ":" + strconv.Itoa(managerPort) + LOG_FILE_DOWNLOAD + "?api-token=" + token;

	for _, logName := range res.Data {
		downloadInfo[logName] = downloadUrl + "&file=" + url.QueryEscape(logName);
	}

	return nil, downloadInfo;
}