package upload

import (
	"os"
	"path"
	"testing"
)

func getRootPath() string {
	pwd, _ := os.Getwd()
	return path.Dir(path.Dir(pwd))
}

func TestCheckPermission(t *testing.T) {

	var filename string = getRootPath() + "/storage/upload/1.txt"

	if !CheckPermission(filename) {
		t.Fatalf("上传目录没有权限")
	}
}

func TestCheckSavePath(t *testing.T) {

	var filename string = getRootPath() + "/storage/upload/1.txt"

	if !CheckSavePath(filename) {
		t.Fatalf("文件不存在")
	}

}