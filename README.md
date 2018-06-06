> ### 最近已将 wmq 和 wmq-admin 进行全部重构，由于重构后代码架构，实现方式等相差较大。故新建项目 [WMQX](https://github.com/phachon/wmqx) 对应的操作后台为 [WMQX-UI](https://github.com/phachon/wmqx-ui)。重构后的 [WMQX](https://github.com/phachon/wmqx) 代码分层清晰，api 更简洁标准，易于维护和扩展。当然，您还可以继续使用 wmq 和 wmq-admin

# Wmq-Admin
[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/phachon/wmq-admin/) 
[![license](https://img.shields.io/github/license/phachon/wmq-admin.svg?style=plastic)]()
----
根据 WMQ 服务提供的后台管理系统
实现了对 WMQ 服务的用户，节点，消息，消费，日志的统一操作管理

## WMQ
基于 RabbitMQ 开发的消息队列服务，支持 http 协议  
fork: https://github.com/snail007/wmq

## 环境依赖
- go1.8
- beego1.8
- sqllite

## 功能
- 公告管理  
及时发布系统公告，通知升级
- 用户管理  
添加用户
- 节点管理  
多台机器多节点部署
- 消息管理  
对消息的增、删、改、查、重载、测试功能
- 消费管理  
对消费者的增、删、改、查、重载功能
- 日志管理  
日志检索和日志下载

# 部署安装
## 安装
下载已经编译好的二进制程序，地址 https://github.com/phachon/wmq-admin/releases  
根据环境变量部署  
设置PATH环境变量  
GOENV = development  
GOENV = production  
GOENV = testing  

解压目录执行 ./wmq-admin  

## 运行
运行 nohup ./wmq-admin &  
浏览器访问 http://127.0.0.1:8080  
登录初始账号密码：root 123456  


## 反馈

欢迎提交意见和代码，联系方式 phachon@163.com

## License

MIT

Thanks
---------
Create By phachon@163.com
