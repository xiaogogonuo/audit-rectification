package handler

import (
	"github.com/gin-gonic/gin"
)

type handler interface {
	Handle(*gin.Engine)
}

type Handler struct {
	Handlers []handler
}

func NewHandle(h ...handler) *Handler {
	return &Handler{Handlers: h}
}

func (hr *Handler) AddHandler(h handler) {
	hr.Handlers = append(hr.Handlers, h)
}
