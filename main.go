package main

import (
	"altair-backend/config"
	_ "altair-backend/config"
	"altair-backend/internal/router"
	"altair-backend/internal/task"
	"altair-backend/log"
	"altair-backend/pkg/validate"
	"net/http"
	"time"
)

//@title 数字资产管理后台
//@version 1.0
//@Description 数字资产管理后台第一版本
func main() {
	err := validate.InitTrans("zh")
	if err != nil {
		log.Fatal(err.Error())
	}
	task.CronInit()
	r := router.NewRouter()
	s := &http.Server{
		Addr:           ":" + config.ServerConfig.HttpPort,
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
