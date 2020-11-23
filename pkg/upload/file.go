package upload

import (
	"fmt"
	"gotour/blog-service/global"
	"gotour/blog-service/pkg/helper"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType uint8

const (
	TypeImage FileType = iota + 1
	TypeExcel
	TYpeTxt
)

// GetFileName 返回Md5加密后的文件名
func GetFileName(file multipart.File, filename string) string {

	ext := GetFileExt(filename)
	content, _ := ioutil.ReadAll(file)
	filename = helper.EncodeMd5(string(content))

	return filename + "." + ext
}

// GetFileExt 获取文件后缀
func GetFileExt(filename string) string {
	return path.Ext(path.Base(filename))[1:]
}

func GetSavePath() string {
	return global.UploadSetting.SavePath
}

// CheckSavePath 检查目标路径是否存在文件， true 表示不存在文件
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// CheckContainExt 检测文件后缀是否允许
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)

	fmt.Println(ext, global.UploadSetting.ImageAllowExts)
	switch t {
	case TypeImage:
		for _, allowExt := range global.UploadSetting.ImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

// CheckMaxSize 检测文件大小是否符合配置要求
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size > global.UploadSetting.ImageMaxSize*1024*1024 {
			return false
		}
	}

	return true
}

// CheckPermission 检测是否有权限创建文件 true 有权限
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return !os.IsPermission(err)
}

func CreateSavePath(path string, fileMode os.FileMode) error {
	err := os.MkdirAll(path, fileMode)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(fileHeader *multipart.FileHeader, dst string) (err error) {

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}

	srcFile, err := fileHeader.Open()
	if err != nil {
		return
	}
	_, err = io.Copy(dstFile, srcFile)
	return
}
