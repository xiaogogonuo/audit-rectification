package rect

import (
	"audit-rectification/internal/rectserver/model"
	"audit-rectification/internal/rectserver/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerRect struct {
	s service.Service
}

func NewHandlerRect(s service.Service) *HandlerRect {
	return &HandlerRect{s}
}

func (hr *HandlerRect) Handle(engine *gin.Engine) {
	engine.POST("/rect", hr.Rect)
}

func (hr *HandlerRect) Rect(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			ErrorMessage: err.Error(),
			Result: nil,
		})
		return
	}
	c.Set("id", 100)
	res, err := hr.s.Fetching(c, req.Content)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusBadRequest,
			ErrorMessage: err.Error(),
			Result: nil,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
