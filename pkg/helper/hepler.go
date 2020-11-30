package helper

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
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

// MultiFormFile 获取多个上传文件
func MultiFormFile(c *gin.Context, key string) ([]*multipart.FileHeader, error) {
	form, err := c.MultipartForm()
	if form != nil && form.File != nil {
		if fhs := form.File[key]; len(fhs) > 0 {
			return fhs, nil
		}
	}
	return nil, err
}

// DelFiles 批量删除文件
func RemoveFile(paths ...string) {
	for _, path := range paths {
		_  = os.Remove(path)
	}
}