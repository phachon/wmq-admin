package common

import "time"

type Date struct{}

//格式化 unix 时间戳
func (date *Date) format(unixTime int64, format string) string {
	return time.Unix(unixTime, 0).Format(format);
}
