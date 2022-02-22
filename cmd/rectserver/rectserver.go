package main

import (
	"audit-rectification/internal/rectserver/api"
	"audit-rectification/internal/rectserver/api/handler"
	"audit-rectification/internal/rectserver/app"
	"audit-rectification/internal/rectserver/repo"
	"audit-rectification/internal/rectserver/service"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
	"time"
)

func main() {
	rep := repo.NewMysqlRepository(nil)
	sev := service.NewRectService(rep)
	handle := handler.NewHandler(sev)
	engine := app.NewGinEngine()
	router := api.NewRouter(handle)
	// 服务注册
	app.Register(engine, router)

	// 不停机重启应用
	endless.DefaultReadTimeOut = 10 * time.Second  // 读超时时间为10s
	endless.DefaultWriteTimeOut = 30 * time.Second // 写超时时间为30s
	endless.DefaultMaxHeaderBytes = 1 << 20        // 请求头最大为1m
	addr := fmt.Sprintf(":%d", 8888)
	srv := endless.NewServer(addr, engine)
	srv.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
