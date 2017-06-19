package common

import (
	"encoding/hex"
	"crypto/md5"
)

//md5加密
func Md5Encode(str string) string {
	hash := md5.New();
	hash.Write([]byte(str));
	return hex.EncodeToString(hash.Sum(nil));
}