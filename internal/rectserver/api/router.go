package api

import (
	"audit-rectification/internal/rectserver/api/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*handler.RectHandler
}

func (r *Router) With(engine *gin.Engine) {
	engine.POST("/rect", r.FetchRect)
}

func NewRouter(rh *handler.RectHandler) *Router {
	return &Router{rh}
}
