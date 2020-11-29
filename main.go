package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/natefinch/lumberjack.v2"
	"gotour/blog-service/global"
	"gotour/blog-service/internal/model"
	"gotour/blog-service/internal/routers"
	"gotour/blog-service/pkg/logger"
	"gotour/blog-service/pkg/setting"
	"gotour/blog-service/pkg/tracer"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	port    string
	runMode string
	config  string

	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {

	// 设置启动参数
	setupFlag()

	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}

	setupLogger()
}

// @title 博客系统
// @version 1.0
// @description Go语言编程练习
// @termsOfService http://fym123.top
func main() {
	if isVersion {
		printVersionInfo()
		return
	}

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		fmt.Println("Start server http://locahost:" + global.ServerSetting.HttpPort)
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServer err: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shuting down server...")

	// 收到退出信号(CTRL + C)后，最长等待5秒
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}

func printVersionInfo() {
	fmt.Printf("build time: %s\n", buildTime)
	fmt.Printf("build version: %s\n", buildVersion)
	fmt.Printf("git commit id: %s\n", gitCommitID)
}

func setupFlag() {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式, debug or release")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "显示编译信息")
	flag.Parse()
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

	mySetting, err := setting.NewSetting(strings.Split(config, ",")...)
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
	global.AppSetting.ContextTimeout *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}

	runModes := map[string]bool{
		"debug":   true,
		"release": true,
	}
	if runMode != "" {
		runMode = strings.ToLower(runMode)
		if _, ok := runModes[runMode]; !ok {
			return errors.New("runMode can only be 'debug' or 'release'.")
		}
		global.ServerSetting.RunMode = runMode
	}

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
		MaxSize:   100,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}
