package api

import (
	"github.com/gin-gonic/gin"
	"gotour/blog-service/global"
	"gotour/blog-service/internal/service"
	"gotour/blog-service/pkg/app"
	"gotour/blog-service/pkg/convert"
	"gotour/blog-service/pkg/errcode"
	"gotour/blog-service/pkg/helper"
	"gotour/blog-service/pkg/upload"
	"path"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

// MultiFormFile 文件批量上传
func (u Upload) MultiFormFile(c *gin.Context) {
	response := app.NewResponse(c)
	fhs, err := helper.MultiFormFile(c, "file")
	fileType := convert.StrTo(c.PostForm("type")).MustUInt8()
	if err != nil {
		errRep := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRep)
		return
	}

	if fileType == 0 || len(fhs) == 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	ser := service.New(c.Request.Context())
	var accessUrls = make([]string, len(fhs))
	var savePaths = make([]string, len(fhs))
	for i, fileHeader := range fhs {
		file, err := fileHeader.Open()
		if err != nil {
			helper.RemoveFile(savePaths...)
			errRep := errcode.InvalidParams.WithDetails(err.Error())
			response.ToErrorResponse(errRep)
			return
		}
		fileInfo, err := ser.UploadFile(upload.FileType(fileType), file, fileHeader)
		if err != nil {
			helper.RemoveFile(savePaths...)
			global.Logger.Errorf(c, "service.UploadFile err: %v", err)
			errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}

		accessUrls[i] = fileInfo.AccessUrl
		savePaths[i] = path.Join(global.UploadSetting.SavePath, fileInfo.Name)
	}

	response.ToResponse(gin.H{
		"file_access_url": accessUrls,
	})
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustUInt8()
	if err != nil {
		errRep := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRep)
		return
	}

	if fileHeader == nil || fileType == 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	ser := service.New(c.Request.Context())
	fileInfo, err := ser.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "service.UploadFile err: %v", err)
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
