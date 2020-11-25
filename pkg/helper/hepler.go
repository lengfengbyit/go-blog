package helper

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// str2md5 字符串 md5 加密
func EncodeMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func TimeFmt(t time.Time, typ string) string {
	var timeFmt string = "2006-01-02 15:04:05"
	switch typ {
	case "date":
		timeFmt = timeFmt[:10]
	case "time":
		timeFmt = timeFmt[11:]
	}
	return t.Format(timeFmt)
}
