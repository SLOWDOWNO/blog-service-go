package main

import (
	"log"
	"net/http"
	"time"

	"github.com/SLOWDOWNO/blog-service-go/global"
	"github.com/SLOWDOWNO/blog-service-go/internal/model"
	"github.com/SLOWDOWNO/blog-service-go/internal/routers"
	"github.com/SLOWDOWNO/blog-service-go/pkg/logger"
	"github.com/SLOWDOWNO/blog-service-go/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 初始化全局变量
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setUpLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
}

// setupSetting函数用于设置全局配置。
// 它会读取配置文件中的Server、App和Database部分，并将其存储到全局变量中。
// 设置全局配置后，会将读取超时和写入超时转换为秒。
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setUpLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
