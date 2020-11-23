package service

import (
	"github.com/pkg/errors"
	"gotour/blog-service/global"
	"gotour/blog-service/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (ser *Service) UploadFile(fileType upload.FileType, file multipart.File,
	header *multipart.FileHeader) (*FileInfo, error) {

	fileName := upload.GetFileName(file, header.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName

	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("File suffix is not supported.")
	}

	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("Failed to created save directory.")
		}
	}

	if !upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("Exceeded maximum file limit.")
	}

	if !upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("Insufficient file permission.")
	}

	if err := upload.SaveFile(header, dst); err != nil {
		return nil, err
	}

	accessUrl := global.UploadSetting.ServerUrl + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessUrl: accessUrl,
	}, nil
}
