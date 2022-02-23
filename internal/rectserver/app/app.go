package app

import (
	"audit-rectification/internal/rectserver/handler"
	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}

// Register 将所有handler注册到engine
func Register(engine *gin.Engine, handler *handler.Handler) {
	for _, hx := range handler.Handlers {
		hx.Handle(engine)
	}
}
