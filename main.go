package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/where/:year", func(c *gin.Context) {
		year, err := strconv.Atoi(c.Param("year"))
		target := time.Date(year, time.January,1,0,0,0,0, time.UTC)
		if err != nil {
			log.Println(err)
			return
		}
		if year < time.Now().Year() {
			daysGone := int64(time.Since(target).Hours()) / 24
			c.String(http.StatusOK, "Days gone: %d", daysGone)
		} else {
			daysLeft := int64(time.Until(target).Hours()) / 24
			c.String(http.StatusOK, "Days left: %d", daysLeft)
		}
	})
	err := router.Run(":3000")
	if err != nil {
		log.Println(err)
		return
	}
}