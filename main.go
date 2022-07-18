package main

import (
	"github.com/YingML/gin-blog-demo/global"
	"github.com/YingML/gin-blog-demo/internal/routers"
	"github.com/YingML/gin-blog-demo/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func setupSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return nil
	}
	err = set.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	log.Printf("%+v", global.ServerSetting)
	log.Printf("%+v", global.AppSetting)
	log.Printf("%+v", global.DatabaseSetting)
	return nil
}

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
	s.ListenAndServe()
}
