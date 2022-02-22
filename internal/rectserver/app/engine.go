package app

import "github.com/gin-gonic/gin"

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}
