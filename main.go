package main

import (
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"log"
	"net/http"
)

func main() {
	app := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        app,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("监听端口失败：%v", err)
	}
}
