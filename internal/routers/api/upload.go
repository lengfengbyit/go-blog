package api

import (
	"github.com/gin-gonic/gin"
	"gotour/blog-service/global"
	"gotour/blog-service/internal/service"
	"gotour/blog-service/pkg/app"
	"gotour/blog-service/pkg/convert"
	"gotour/blog-service/pkg/errcode"
	"gotour/blog-service/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
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
		global.Logger.Errorf("service.UploadFile err: %v", err)
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
