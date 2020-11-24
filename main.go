package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"gotour/blog-service/global"
	"gotour/blog-service/internal/model"
	"gotour/blog-service/internal/routers"
	"gotour/blog-service/pkg/logger"
	"gotour/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	setupLogger()
}

func setupSetting() error {
	settingMap := map[string]interface{}{
		"Server":   &global.ServerSetting,
		"App":      &global.AppSetting,
		"Database": &global.DatabaseSetting,
		"upload":   &global.UploadSetting,
		"JWT":      &global.JWTSetting,
		"Email":    &global.EmailSetting,
	}

	mySetting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	for key, val := range settingMap {
		err = mySetting.ReadSection(key, val)
		if err != nil {
			return err
		}
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	// 这里不能使用 := 来初始化 global.DBEngine
	// 因为 := 会重新创建左侧变量，导致并没有初始化global包中的DBEngine变量
	// 在当前包可以使用global.DBEngine，在其他包中就不能使用了
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() {

	dateSuffix := ""
	if global.AppSetting.LogDateSuffix {
		dateSuffix = "." + time.Now().Format("20060102")
	}

	filePath := fmt.Sprintf("%s/%s%s%s",
		global.AppSetting.LogSavePath,
		global.AppSetting.LogFileName,
		dateSuffix,
		global.AppSetting.LogFileExt)

	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  filePath,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
}

// @title 博客系统
// @version 1.0
// @description Go语言编程练习
// @termsOfService http://fym123.top
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Start server http://locahost:" + global.ServerSetting.HttpPort)
	_ = s.ListenAndServe()
}
