package server

import (
	"github.com/gin-gonic/gin"
	"qsoft/internal/handler"
	"qsoft/internal/middleware"
)

type Server struct {
	h *handler.YearHandler
}

func NewServer(h *handler.YearHandler) *Server {
	return &Server{h: h}
}

func (s *Server) Run() error {
	router := gin.Default()
	router.Use(middleware.PingMiddleware())
	router.GET("/where/:year", s.h.Where)
	err := router.Run(":3000")
	if err != nil {
		return err
	}
	return nil
}
