package helper

import (
	"crypto/md5"
	"encoding/hex"
)

// str2md5 字符串 md5 加密
func EncodeMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}


