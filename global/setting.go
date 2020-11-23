package global

import (
	"gotour/blog-service/pkg/logger"
	"gotour/blog-service/pkg/setting"
)

// 全局配置

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	UploadSetting   *setting.UploadSetting
	JWTSetting      *setting.JWTSetting
	EmailSetting    *setting.EmailSetting
	Logger          *logger.Logger
)
