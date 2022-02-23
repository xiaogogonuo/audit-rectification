package service

import (
	"audit-rectification/internal/rectserver/model"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Fetching(*gin.Context, string) (model.Response, error)
}
