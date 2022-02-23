package main

import (
	"audit-rectification/internal/rectserver/app"
	"audit-rectification/internal/rectserver/handler"
	rh "audit-rectification/internal/rectserver/handler/rect"
	"audit-rectification/internal/rectserver/initialize"
	rr "audit-rectification/internal/rectserver/repo/rect"
	rs "audit-rectification/internal/rectserver/service/rect"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
	"time"
)

func main() {
	// 初始化nebula图数据库
	nebula := initialize.NebulaInit()
	// 新建nebula仓库
	nebulaRepository := rr.NewNebulaRepository(nebula)
	// 新建整改服务
	rectService := rs.NewServiceRect(nebulaRepository)
	// 新建整改handler
	rectHandler := rh.NewHandlerRect(rectService)
	// 新建全局handler
	handlers := handler.NewHandle()
	// 将整改handler添加到全局handler
	handlers.AddHandler(rectHandler)
	// 新建gin引擎
	engine := app.NewEngine()
	// 将所有handler注册到gin引擎
	app.Register(engine, handlers)

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
