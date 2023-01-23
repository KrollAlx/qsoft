package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"qsoft/internal/service"
	"strconv"
)

type YearHandler struct {
	s service.Year
}

func NewYearHandler(s service.Year) *YearHandler {
	return &YearHandler{s: s}
}

func (h *YearHandler) Where(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		log.Println(err)
		return
	}
	result := h.s.WhereYear(year)
	c.String(http.StatusOK, result)
}
