package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/orm"
)

//初始化
func Init() {
	databaseType := beego.AppConfig.String("database.type");
	if(databaseType == "mysql") {
		mysqlConn();
	}
	if(databaseType == "sqlite") {
		sqliteConn();
	}

	orm.RegisterModel(new(User), new(Node), new(Notice));

	if beego.AppConfig.String("runmode") == "development" {
		orm.Debug = true
	}
}

//mysql 连接
func mysqlConn()  {
	host := beego.AppConfig.String("database.mysql.host");
	port := beego.AppConfig.String("database.mysql.port");
	user := beego.AppConfig.String("database.mysql.user");
	password := beego.AppConfig.String("database.mysql.password");
	database := beego.AppConfig.String("database.mysql.name");
	charset := beego.AppConfig.String("database.mysql.charset");
	maxIdle := beego.AppConfig.DefaultInt("database.mysql.maxIdle", 30);
	maxOpenConn := beego.AppConfig.DefaultInt("database.mysql.maxOpenConn", 30);
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset;

	//驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册
	orm.RegisterDataBase("default", "mysql", dsn);
	//数据库最大空闲连接
	orm.SetMaxIdleConns("default", maxIdle);
	//数据库最大连接
	orm.SetMaxOpenConns("default", maxOpenConn);
}

//sqlite 连接
func sqliteConn()  {
	sqlite := beego.AppConfig.String("database.sqlite.path");
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", sqlite)
}

func TableName(name string) string {
	return beego.AppConfig.String("database.prefix") + name;
}