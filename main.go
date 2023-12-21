package main

import (
	"log"
	"net/http"
	"time"

	"github.com/SLOWDOWNO/blog-service-go/global"
	"github.com/SLOWDOWNO/blog-service-go/internal/model"
	"github.com/SLOWDOWNO/blog-service-go/internal/routers"
	"github.com/SLOWDOWNO/blog-service-go/pkg/setting"
)

// 初始化全局变量
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatal("init.setupDBEngine err: %v", err)
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
