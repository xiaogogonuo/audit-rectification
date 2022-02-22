package app

import (
	"audit-rectification/internal/rectserver/api"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine, router *api.Router) {
	router.With(engine)
}
