package handler

import (
	"audit-rectification/internal/rectserver/model"
	"audit-rectification/internal/rectserver/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RectHandler struct {
	ser service.Service
}

func NewHandler(s service.Service) *RectHandler {
	return &RectHandler{s}
}

func (r *RectHandler) FetchRect(c *gin.Context) {
	var request model.Request
	// 入参检测
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: err.Error(),
			Result:       nil,
		})
		return
	}
	// 业务查询
	res, err := r.ser.Fetch(c, request.Content)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: err.Error(),
			Result:       nil,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
