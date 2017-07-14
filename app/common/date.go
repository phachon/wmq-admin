package common

import (
	"github.com/astaxie/beego"
	"time"
)

type Date struct{}

//格式化 unix 时间戳
func (date *Date) Format(unixTime int64, format string) string {
	return beego.Date(time.Unix(unixTime, 0), format);
}
