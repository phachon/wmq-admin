package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"wmq-admin/app/common"
)

const (
	USER_IS_DELETE_TRUE = 1
	USER_IS_DELETE_FALSE = 0
);

type User struct {
	Id int64 `orm:"auto;pk;column(user_id);"`
	Name string
	Email string
	Password string
	Mobile string
	IsDelete int
	CreateTime int64
	UpdateTime int64
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

//插入一条用户信息
func InsertUser(user *User) (int64, error){

	if(user.Name == "") {
		return 0, fmt.Errorf("用户名不能为空!");
	}
	if(user.Email == "") {
		return 0, fmt.Errorf("邮箱不能为空!");
	}
	if(user.Password == "") {
		return 0, fmt.Errorf("密码不能为空!");
	}

	user.Password = common.Md5Encode(user.Password);

	user.IsDelete = USER_IS_DELETE_FALSE;
	user.CreateTime = time.Now().Unix();
	user.UpdateTime = time.Now().Unix();

	return orm.NewOrm().Insert(user);
}

//根据 name 查找 user
func GetUserByName(name string) ([]*User)  {
	users := make([]*User, 0)
	query := orm.NewOrm().QueryTable(TableName("user")).Filter("name", name);
	query.All(&users);
	return users;
}

//根据 user_id 查找 user
func GetUserByUserId(userId int) ([]*User)  {
	users := make([]*User, 0)
	query := orm.NewOrm().QueryTable(TableName("user")).Filter("user_id", userId);
	query.All(&users);
	return users;
}

//来修改用户
func UpdateUser(user *User, fields ...string) (int64, error) {
	return orm.NewOrm().Update(user, fields...);
}