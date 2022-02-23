package repo

import (
	"audit-rectification/internal/rectserver/model"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	Fetch(*gin.Context, string) (model.Response, error)
}
