> ### 最近已将 wmq 和 wmq-admin 进行全部重构，由于重构后再代码架构，实现等相差较大。故新建项目 [WMQX](https://github.com/phachon/wmqx) 对应的操作后台为 [WMQX-UI](https://github.com/phachon/wmqx-ui)。重构后的 [WMQX](https://github.com/phachon/wmqx) 代码分层清晰，api 更简洁标准，易于维护和扩展。当然，您还可以继续使用 wmq 和 wmq-admin

# Wmq-Admin
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
下载最新版本的源代码  
根据环境变量部署  
设置 go 环境变量  
GOENV = development  
GOENV = production  
GOENV = testing  

## 编译

- 手动编译  
搭建安装Go环境，进入项目根目录  
执行：go get ./...  
编译：go build ./  

- 免编译  
下载已经编译好的二进制程序，地址 https://github.com/phachon/wmq-admin-release  
将下载好的文件重命名为 wmq-admin, 并放置在项目根目录  

## 运行
运行 nohup ./wmq-admin &  
登录初始账号密码：root 123456  


## 反馈

欢迎提交意见和代码，联系方式 phachon@163.com

## License

MIT

Thanks
---------
Create By phachon@163.com
