package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id int `orm:"auto;pk;column(user_id);"`
	Name string
	Email string
	Password string
	Mobile string
	Status int
	CreateTime int
	UpdateTime int
}

//table name
func (user *User) TableName() string {
	return TableName("user");
}

//获取所有的 user
func GetUsers() ([]*User) {
	users := make([]*User, 0)
	query := orm.NewOrm().QueryTable(TableName("user"))
	query.All(&users);
	return users;
}