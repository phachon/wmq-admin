package common

import (
	"encoding/base64"
	"strings"
	"crypto/md5"
	"encoding/hex"
)

const (
	BASE_64_TABLE        = "<,./?~!@#$CDVWX%^&*ABYZabcghijkpqrstuvwxyz01EFKLMNOPQRSTU2345678"
	HASH_FUNCTION_HEADER = "zh.ife.iya"
	HASH_FUNCTION_FOOTER = "09.O25.O20.78"
)

type Encrypt struct {}

//base64 加密
func (encrypt *Encrypt) Base64Encode(str string) string {
	var coder = base64.NewEncoding(BASE_64_TABLE)
	var src []byte = []byte(HASH_FUNCTION_HEADER + str + HASH_FUNCTION_FOOTER)
	return string([]byte(coder.EncodeToString(src)))
}

//base64 解密
func (encrypt *Encrypt) Base64Decode(str string) (string, error) {
	var src []byte = []byte(str)
	var coder = base64.NewEncoding(BASE_64_TABLE)
	by, err := coder.DecodeString(string(src))
	return strings.Replace(strings.Replace(string(by), HASH_FUNCTION_HEADER, "", -1), HASH_FUNCTION_FOOTER, "", -1), err
}

//md5加密
func (encrypt *Encrypt) Md5Encode(str string) string {
	hash := md5.New();
	hash.Write([]byte(str));
	return hex.EncodeToString(hash.Sum(nil));
}