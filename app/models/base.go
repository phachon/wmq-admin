package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

//初始化
func Init()  {
	host := beego.AppConfig.String("database.host");
	port := beego.AppConfig.String("database.port");
	user := beego.AppConfig.String("database.user");
	password := beego.AppConfig.String("database.password");
	database := beego.AppConfig.String("database.name");
	charset := beego.AppConfig.String("database.charset");
	maxIdle := beego.AppConfig.DefaultInt("database.maxIdle", 30);
	maxOpenConn := beego.AppConfig.DefaultInt("database.maxOpenConn", 30);
	//timezone := beego.AppConfig.String("database.timezone");

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset;

	//驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册
	orm.RegisterDataBase("default", "mysql", dsn);
	//数据库最大空闲连接
	orm.SetMaxIdleConns("default", maxIdle);
	//数据库最大连接
	orm.SetMaxOpenConns("default", maxOpenConn);
	//时区设置
	//orm.DefaultTimeLoc = timezone

	orm.RegisterModel(new(User))

	if beego.AppConfig.String("runmode") == "development" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("database.prefix") + name;
}